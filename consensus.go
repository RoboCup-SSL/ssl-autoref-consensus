package main

import (
	"flag"
	. "github.com/RoboCup-SSL/ssl-go-tools/pkg/sslproto"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"io"
	"log"
	"math"
	"net"
	"time"
)

var majorityTimeout = flag.Int64("mt", 9, "Majority timeout in seconds")
var replyTimeout = flag.Int64("rt", 10, "Reply timeout in seconds")

var numClients = 0
var messageId = uint32(0)

func main() {
	listeningAddr := flag.String("l", ":10008", "The listening address where the consensus is expecting requests from autoRefs")
	refBoxAddr := flag.String("r", "localhost:10007", "The address of the ssl-refbox")
	flag.Parse()

	listener, err := net.Listen("tcp", *listeningAddr)
	if err != nil {
		log.Fatalf("could not listen on %v", *listeningAddr)
	}
	defer listener.Close()
	log.Printf("Listening on %s", *listeningAddr)

	refBoxConn, err := net.Dial("tcp", *refBoxAddr)
	if err != nil {
		log.Printf("could not connect to refbox at %v", *refBoxAddr)
		refBoxConn = nil
	} else {
		defer refBoxConn.Close()
		log.Printf("Connected to refbox at %v", *refBoxAddr)
	}

	controlRequests := make(chan ASyncRequest)

	go handleConsensus(controlRequests, refBoxConn)

	for {
		if conn, err := listener.Accept(); err == nil {
			go handleClientConnection(conn, controlRequests)
		} else {
			continue
		}
	}
}

func handleConsensus(requests <-chan ASyncRequest, refBoxConn net.Conn) {

	reqBuffer := make([]*ASyncRequest, 0)

	for {
		select {
		case aSyncRequest := <-requests:
			aSyncRequest.replied = false

			reqBuffer = append(reqBuffer, &aSyncRequest)
			reqBuffer = removeTimedOutRequests(reqBuffer)

			reqBuffer = replyOnMajority(reqBuffer, aSyncRequest.request, refBoxConn)
		case <-time.After(time.Millisecond * 100):
			reqBuffer = removeTimedOutRequests(reqBuffer)
		}
	}
}

func replyOnMajority(reqBuffer []*ASyncRequest, request *SSL_RefereeRemoteControlRequest, refBoxConn net.Conn) []*ASyncRequest {

	matchingRequests := findMatchingRequests(reqBuffer, request)
	numMatching := len(matchingRequests)
	majority := int(math.Floor(float64(numClients) / 2.0))

	if numMatching > majority {
		log.Printf("Found majority: %d/%d", numMatching, numClients)

		refboxRequest := SSL_RefereeRemoteControlRequest{}
		proto.Merge(&refboxRequest, request)
		*refboxRequest.MessageId = messageId
		messageId++

		var outcome SSL_RefereeRemoteControlReply_Outcome
		if refBoxConn == nil {
			outcome = SSL_RefereeRemoteControlReply_OK
		} else if err := SendMessage(refBoxConn, &refboxRequest); err != nil {
			log.Println("unable to send reply to refbox", err)
			outcome = SSL_RefereeRemoteControlReply_COMMUNICATION_FAILED
		} else {
			reply := new(SSL_RefereeRemoteControlReply)
			if err := ReceiveMessage(refBoxConn, reply); err == nil {
				outcome = reply.GetOutcome()
			} else {
				log.Println("unable to receive reply", err)
				outcome = SSL_RefereeRemoteControlReply_COMMUNICATION_FAILED
			}
		}

		for _, req := range matchingRequests {
			log.Println("reply: ", outcome)
			req.Reply(outcome)
		}

		var remainingRequests []*ASyncRequest
		for _, req := range reqBuffer {
			if !req.replied {
				remainingRequests = append(remainingRequests, req)
			}
		}
		return remainingRequests
	}
	return reqBuffer
}

func removeTimedOutRequests(reqBuffer []*ASyncRequest) []*ASyncRequest {
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
	return reqBuffer
}

func handleClientConnection(clientConn net.Conn, messages chan<- ASyncRequest) {
	numClients++
	log.Printf("Connection established: %v, now %d clients", clientConn.RemoteAddr(), numClients)

	// Close the connection when the function exits
	defer clientConn.Close()

	for {
		err := handleClientRequest(clientConn, messages)
		if errors.Cause(err) == io.EOF {
			// connection is closed
			break
		}
		if err != nil {
			log.Println("unable to handle client request: ", err)
		}
	}

	numClients--
	log.Printf("Connection closed: %v, now %d clients", clientConn.RemoteAddr(), numClients)
}

func handleClientRequest(clientConnection net.Conn, requests chan<- ASyncRequest) error {

	request := new(SSL_RefereeRemoteControlRequest)
	if err := ReceiveMessage(clientConnection, request); err != nil {
		return errors.Wrap(err, "unable to receive request from client")
	}

	log.Println("Received message:", request)

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
	case <-time.After(time.Duration(*replyTimeout) * time.Second):
		log.Printf("No reply received")
		outcome = SSL_RefereeRemoteControlReply_COMMUNICATION_FAILED
	}

	reply := &SSL_RefereeRemoteControlReply{
		MessageId: request.MessageId,
		Outcome:   &outcome,
	}

	if err := SendMessage(clientConnection, reply); err != nil {
		return errors.Wrap(err, "unable to send reply to client")
	}
	return nil
}
