package connection

import (
	"net"
	"gobuff/src/com/gobuff/transfer"
	"gobuff/src/com/gobuff/constant"
	pb "gobuff/src/com/gobuff/proto"
	"fmt"
	"time"
	"encoding/binary"
	"gobuff/src/com/gobuff/heartbeat"
	"gobuff/src/com/gobuff/cache"
	"gobuff/src/com/gobuff/log"
)

func NewConnection(conn net.Conn) {
	stop := make(chan bool)
	send := make(chan bool)
	go heartbeat.ServerDeal(conn, stop, send)
	handleRec(conn, stop, send)
}

func handleSend(conn net.Conn) {
	for {
		t := time.Now()
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(t.UnixNano()))
		data := &pb.Data{
			Code: 888,
			Uid:  "uid",
			Data: b,
		}
		err := transfer.Write(conn, data)
		if err != nil {
			panic(err)
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func handleRec(conn net.Conn, stop chan bool, send chan bool) {
L:
	for {
		select {
		case <-stop:
			delete(cache.Clients, conn.RemoteAddr().String())
			log.Log("连接超时断开", conn.RemoteAddr().String())
			break L
		default:
			{
				rec := &pb.Data{}
				err := transfer.Read(conn, rec)
				if err != nil {
					panic(err)
				}
				if rec.Code == constant.CodeHeartBeat {
					fmt.Println("收到客户端心跳包")
					heartbeat.SendHeartBeat(conn)
					send <- true
				} else if rec.Code == constant.CodeUserId {
					cache.Clients[rec.Uid] = conn
					fmt.Println(cache.Clients)
				} else {
					//fmt.Println("接收到数据：", conn.RemoteAddr(), rec)
				}

				time.Sleep(50 * time.Millisecond)
			}

		}

	}

}
