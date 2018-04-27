package transfer

import (
	"sync"
	"net"
	"github.com/golang/protobuf/proto"
	"gobuff/src/com/gobuff/serialize"
)

var mutex sync.Mutex

func Write(conn net.Conn, message proto.Message) error {
	buff, err := serialize.ToBytes(message)
	if err != nil {
		return err
	}
	conn.Write(buff)
	return nil
}

func Read(conn net.Conn, message proto.Message) error {
	return serialize.ToProto(conn, message)
}
