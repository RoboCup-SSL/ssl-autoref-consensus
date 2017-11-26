package main

import (
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"math"
	"net"
	"time"
)

const majorityTimeout = time.Second * 9
const replyTimeout = time.Second * 10
const positionToleranceMM = 100

var numClients = 0

type ASyncRequest struct {
	request     *SSL_RefereeRemoteControlRequest
	outcome     chan<- SSL_RefereeRemoteControlReply_Outcome
	receiveTime time.Time
	replied     bool
}

func (r *ASyncRequest) Reply(outcome SSL_RefereeRemoteControlReply_Outcome) {
	if !r.replied {
		r.outcome <- outcome
		r.replied = true
	}
}

func (r *ASyncRequest) TimedOut() bool {
	return r.receiveTime.Add(majorityTimeout).Before(time.Now())
}

func main() {
	addr := ":10007"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("could not connect to %v", addr)
	}
	log.Printf("Listening on %s", addr)

	controlRequests := make(chan ASyncRequest)

	go handleConsensus(controlRequests)

	for {
		if conn, err := ln.Accept(); err == nil {
			go handleProtoClient(conn, controlRequests)
		} else {
			continue
		}
	}
}

func handleConsensus(requests <-chan ASyncRequest) {

	reqBuffer := make([]*ASyncRequest, 0)

	for {
		aSyncRequest := <-requests
		request := aSyncRequest.request
		reqBuffer = append(reqBuffer, &aSyncRequest)
		aSyncRequest.replied = false

		numTimedOut := 0
		for _, req := range reqBuffer {
			if req.TimedOut() {
				log.Printf("Request timed out: %v", req.request)
				req.Reply(SSL_RefereeRemoteControlReply_NO_MAJORITY)
				numTimedOut++
			} else {
				break
			}
		}
		reqBuffer = reqBuffer[numTimedOut:]

		matchingRequests := matchRequests(reqBuffer, request)

		numMatching := len(matchingRequests)
		majority := int(math.Floor(float64(numClients) / 2.0))
		log.Printf("matching: %d", numMatching)
		if numMatching > majority {
			log.Printf("Found majority: %d/%d", numMatching, numClients)
			// TODO actually sent the request out

			for _, req := range matchingRequests {
				req.Reply(SSL_RefereeRemoteControlReply_OK)
			}
		}
	}
}

func matchRequests(reqBuffer []*ASyncRequest, request *SSL_RefereeRemoteControlRequest) []*ASyncRequest {
	matchingRequests := make([]*ASyncRequest, 0)
	for _, otherRequest := range reqBuffer {
		if otherRequest.request.equals(request) {
			matchingRequests = append(matchingRequests, otherRequest)
		}
	}
	return matchingRequests
}

func handleProtoClient(conn net.Conn, messages chan<- ASyncRequest) {
	log.Printf("Connection established: %v", conn.RemoteAddr())
	numClients++
	log.Printf("Now %d clients", numClients)

	//Close the connection when the function exits
	defer conn.Close()

	for {
		err := handleRequest(conn, messages)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error: ", err)
		}
	}

	numClients--
	log.Printf("Connection closed: %v", conn.RemoteAddr())
}

func handleRequest(conn net.Conn, requests chan<- ASyncRequest) (err error) {
	dataLength, err := readDataLength(conn)
	if err != nil {
		return
	}

	data := make([]byte, dataLength)
	n, err := conn.Read(data)
	if err != nil || n != int(dataLength) {
		return
	}

	request := new(SSL_RefereeRemoteControlRequest)
	if err = proto.Unmarshal(data[0:n], request); err != nil {
		return
	}
	log.Println("Cmd: " + request.Command.String())

	cOutcome := make(chan SSL_RefereeRemoteControlReply_Outcome)
	aSyncRequest := ASyncRequest{
		request:     request,
		outcome:     cOutcome,
		receiveTime: time.Now(),
	}
	requests <- aSyncRequest

	var outcome SSL_RefereeRemoteControlReply_Outcome
	select {
	case res := <-cOutcome:
		outcome = res
	case <-time.After(replyTimeout):
		log.Printf("No reply received")
		outcome = SSL_RefereeRemoteControlReply_NO_MAJORITY
	}

	reply := &SSL_RefereeRemoteControlReply{
		MessageId: request.MessageId,
		Outcome:   &outcome,
	}

	data, err = proto.Marshal(reply)
	if err != nil {
		log.Println("marshaling error: ", err)
		return
	}

	if err = writeDataLength(conn, len(data)); err != nil {
		return
	}
	n, err = conn.Write(data)
	if err != nil || n != len(data) {
		return
	}
	return
}

func readDataLength(conn net.Conn) (length uint32, err error) {
	// The header is a 4 byte big endian uint32
	header := make([]byte, 4)
	h, err := conn.Read(header)
	if err == nil && h == 4 {
		length = binary.BigEndian.Uint32(header)
	} else {
		length = 0
	}
	return
}

func writeDataLength(conn net.Conn, dataLength int) (err error) {
	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(dataLength))
	n, err := conn.Write(header)
	if n != 4 {
		err = errors.New("invalid size written")
	}
	return
}

func (m *SSL_RefereeRemoteControlRequest) equals(o *SSL_RefereeRemoteControlRequest) bool {

	if m.LastCommandCounter != nil || o.LastCommandCounter != nil {
		if m.LastCommandCounter == nil || o.LastCommandCounter == nil || m.GetLastCommandCounter() != o.GetLastCommandCounter() {
			return false
		}
	}

	if m.Stage != nil || o.Stage != nil {
		if m.Stage == nil || o.Stage == nil || m.GetStage() != o.GetStage() {
			return false
		}
	}

	if m.Command != nil || o.Command != nil {
		if m.Command == nil || o.Command == nil || m.GetCommand() != o.GetCommand() {
			return false
		}
	}

	if m.Card != nil || o.Card != nil {
		if m.Card == nil || o.Card == nil || !m.GetCard().equals(o.GetCard()) {
			return false
		}
	}

	if m.DesignatedPosition != nil || o.DesignatedPosition != nil {
		if m.DesignatedPosition == nil || o.DesignatedPosition == nil || m.GetDesignatedPosition().similar(o.GetDesignatedPosition()) {
			return false
		}
	}

	return true
}

func (m *SSL_RefereeRemoteControlRequest_CardInfo) equals(o *SSL_RefereeRemoteControlRequest_CardInfo) bool {
	if o == nil {
		return false
	}
	if m.Team == nil || o.Team == nil || m.Type == nil || o.Type == nil {
		return false
	}
	if m.GetType() != o.GetType() || m.GetTeam() != o.GetTeam() {
		return false
	}
	return true
}

func (m *SSL_Referee_Point) similar(o *SSL_Referee_Point) bool {
	if o == nil {
		return false
	}

	dx := m.GetX() - o.GetX()
	dy := m.GetY() - o.GetY()
	diff := math.Sqrt(float64(dx*dx + dy*dy))

	if diff > positionToleranceMM {
		return false
	}
	return true
}
