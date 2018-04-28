package consumer

import (
	"net"
	"time"
)

type Consumer struct {
	Conn        net.Conn
	Uid         string
	ConnectTime int64
}

func NewConsumer(conn net.Conn, uid string) Consumer {
	consumer := Consumer{Conn: conn, Uid: uid, ConnectTime: time.Now().UnixNano()}
	return consumer
}
