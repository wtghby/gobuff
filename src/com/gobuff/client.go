package main

import (
	"net"
	"fmt"
	"os"
	pb "./proto"
	"github.com/golang/protobuf/proto"
	"time"
	"encoding/binary"
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
	//index := 0
	for {
		send(conn)
		//se := "客户端发送的数据--" + strconv.Itoa(index)
		//conn.Write([]byte(se))
		//index++
		time.Sleep(50)
	}
}

func recServer(conn net.Conn) {
	for {
		//buff := make([]byte, 1024*2, 1024*2)
		//len, err := conn.Read(buff)
		//if err != nil {
		//	fmt.Println("读取数据失败")
		//}
		//if len > 0 {
		//	fmt.Println("[收到消息]：", string(buff[:len]))
		//}
		handleRec(conn)
		//time.Sleep(50)
	}
}

func handleRec(conn net.Conn) {
	buff := make([]byte, 1024*2, 1024*2)

	for {
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "connection error:", err)
			return
		}

		rec := &pb.Data{}
		data := buff[:n]
		err = proto.Unmarshal(data, rec)
		if err != nil {
			panic(err)
		}
		ti := int64(binary.BigEndian.Uint64(rec.Data))
		fmt.Println("接收到数据：", conn.RemoteAddr(), rec)
		fmt.Println("Send Time：", ti)
		now := time.Now()
		fmt.Println("接收时间：", now.UnixNano())

		//send, err := proto.Marshal(rec)
		//if err != nil {
		//	panic(err)
		//}

		//fmt.Println(send)
		//conn.Write(send)
		//fmt.Println("Server send ovwr")
	}
}

func send(conn net.Conn) {
	d := "this is post message"

	data := &pb.Data{
		Code: 22,
		Uid:  "uid",
		Data: []byte(d),
	}
	pData, err := proto.Marshal(data)
	if err != nil {
		panic(err)
	}

	conn.Write(pData)
}
