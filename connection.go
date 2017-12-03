package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"io"
	"net"
)

func sendMessage(conn net.Conn, message proto.Message) error {

	data, err := proto.Marshal(message)
	if err != nil {
		return errors.Wrap(err, "marshaling error")
	}

	if err = writeDataLength(conn, len(data)); err != nil {
		return errors.Wrap(err, "unable to write data length")
	}
	if _, err = conn.Write(data); err != nil {
		return errors.Wrap(err, "unable to write data")
	}

	return nil
}

func receiveMessage(conn net.Conn, message proto.Message) error {

	dataLength, err := readDataLength(conn)
	if err != nil {
		return errors.Wrap(err, "unable to read data length")
	}

	data := make([]byte, dataLength)
	if _, err = io.ReadFull(conn, data); err != nil {
		return errors.Wrap(err, "unable to read data")
	}

	if err = proto.Unmarshal(data, message); err != nil {
		return errors.Wrap(err, "unable to unmarshal data")
	}
	return nil
}
