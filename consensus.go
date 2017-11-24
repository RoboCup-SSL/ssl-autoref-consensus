package main

import (
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"net"
)

func main() {
	log.Println("Started.")
	ln, err := net.Listen("tcp", ":10007")
	if err != nil {
		log.Fatalln("could not connect")
	}
	for {
		if conn, err := ln.Accept(); err == nil {
			go handleProtoClient(conn)
		} else {
			continue
		}
	}
}

func handleProtoClient(conn net.Conn) {
	log.Println("Connection established")

	//Close the connection when the function exits
	defer conn.Close()

	for {
		err := handleRequest(conn)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error: ", err)
		}
	}

	log.Println("Connection closed")
}

func handleRequest(conn net.Conn) (err error) {
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

	status := SSL_RefereeRemoteControlReply_OK
	reply := &SSL_RefereeRemoteControlReply{
		MessageId: request.MessageId,
		Outcome:   &status,
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
