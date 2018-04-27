package server

import (
	"net"
	"gobuff/src/com/gobuff/config"
	"strconv"
	"gobuff/src/com/gobuff/log"
	"time"
	"fmt"
	pb "gobuff/src/com/gobuff/proto"
	"encoding/binary"
	"gobuff/src/com/gobuff/transfer"
	"gobuff/src/com/gobuff/constant"
)

var clients = make(map[string]net.Conn)

func Run(con config.Config) {
	addr := con.ServerIp + ":" + strconv.Itoa(con.Port)
	listener, err := net.Listen(con.Protocol, addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Log("服务器启动成功，等待连接")

	go recLoop()
	go sendLoop()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		log.Log("新连接：", conn.RemoteAddr())

		clients[conn.RemoteAddr().String()] = conn
	}
}

func sendLoop() {
	for {
		for _, conn := range clients {
			handleSend(conn)
		}
		time.Sleep(50 * time.Millisecond)
	}

}

func handleSend(conn net.Conn) {
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
}

func recLoop() {
	for {
		for _, conn := range clients {
			//hrec(conn)
			handleRec(conn)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func handleRec(conn net.Conn) {
	rec := &pb.Data{}
	err := transfer.Read(conn, rec)
	if err != nil {
		panic(err)
	}
	if rec.Code == constant.CodeHeartBeat {
		fmt.Println("收到客户端心跳包")
		sendHeartBeat(conn)
	} else {
		//fmt.Println("接收到数据：", conn.RemoteAddr(), rec)
	}
}

func sendHeartBeat(conn net.Conn) {
	heartBeat := &pb.Data{Code: constant.CodeHeartBeat}
	err := transfer.Write(conn, heartBeat)
	if err != nil {
		panic(err)
	}
	fmt.Println("发送心跳包")
}
