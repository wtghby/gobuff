package heartbeat

import (
	"net"
	"gobuff/src/com/gobuff/constant"
	"gobuff/src/com/gobuff/transfer"
	pb "gobuff/src/com/gobuff/proto"
	"fmt"
	"time"
)

func SendLoop(conn net.Conn) {
	tick := time.NewTicker(constant.HEART_BEAT_PERIOD * time.Second)
	for {
		select {
		case <-tick.C:
			Send(conn)
		}
	}

}

func Send(conn net.Conn) {
	heartBeat := &pb.Data{Code: constant.CodeHeartBeat}
	err := transfer.Write(conn, heartBeat)
	if err != nil {
		panic(err)
	}
	fmt.Println("发送心跳包")
}

func ServerDeal(conn net.Conn) {
	tick := time.NewTicker(constant.HEART_BEAT_PERIOD * time.Second)
	select {}
}
