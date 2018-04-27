package main

import (
	"net"
	"fmt"
	"os"
	pb "gobuff/src/com/gobuff/proto"
	"time"
	"encoding/binary"
	"gobuff/src/com/gobuff/serialize"
)

var ch = make(chan int)

func main() {
	server := "127.0.0.1:8545"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	defer conn.Close()
	go sendStr(conn)
	go recServer(conn)
	//send(conn)
	<-ch
}

func sendStr(conn net.Conn) {
	for {
		send(conn)
		//s(conn)
		time.Sleep(50 * time.Millisecond)
	}
}

func recServer(conn net.Conn) {
	for {
		handleRec(conn)
		time.Sleep(50 * time.Millisecond)
	}
}

func handleRec(conn net.Conn) {
	rec := &pb.Data{}
	err := serialize.ToProto(conn, rec)
	if err != nil {
		panic(err)
	}
	fmt.Println("接收到数据：", conn.RemoteAddr(), rec)
	ti := int64(binary.BigEndian.Uint64(rec.Data))
	fmt.Println("Send Time：", ti)
	now := time.Now()
	fmt.Println("接收时间：", now.UnixNano())
}

func send(conn net.Conn) {
	d := "this is post message"

	data := &pb.Data{
		Code: 22,
		Uid:  "uid",
		Data: []byte(d),
	}
	buff, err := serialize.ToBytes(data)
	if err != nil {
		panic(err)
	}
	conn.Write(buff)
	fmt.Println("send to server,size : ", len(buff))
}
