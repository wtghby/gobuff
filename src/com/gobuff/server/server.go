package server

import (
	"net"
	"../config"
	"strconv"
	"../log"
	"time"
	"github.com/golang/protobuf/proto"
	"fmt"
	pb "../proto"
	"encoding/binary"
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

		handleConn(conn)
	}
}

func sendLoop() {
	//index := 0
	for {
		//str := "当前连接数：" + strconv.Itoa(len(clients)) + "发送次数：" + strconv.Itoa(index)
		for _, conn := range clients {
			handleSend(conn)
		}
		//index++

		time.Sleep(100)
	}

}

func handleSend(conn net.Conn) {
	//d := "this is server message"
	t := time.Now()
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(t.UnixNano()))
	data := &pb.Data{
		Code: 888,
		Uid:  "uid",
		Data: b,
	}
	pData, err := proto.Marshal(data)
	if err != nil {
		panic(err)
	}

	conn.Write(pData)
}

func recLoop() {
	for {
		for _, conn := range clients {
			handleRec(conn)
			//buff := make([]byte, 1024*2, 1024*2)
			//len, err := conn.Read(buff)
			//if err != nil {
			//	log.Loge("读取数据失败")
			//}
			//if len > 0 {
			//	log.Log("[收到消息]：", string(buff[:len]))
			//}
		}
		time.Sleep(100)
	}
}

func handleRec(conn net.Conn) {
	buff := make([]byte, 1024*2, 1024*2)

	for {
		n, err := conn.Read(buff)
		if err != nil {
			log.Log(conn.RemoteAddr().String(), "connection error:", err)
			return
		}

		rec := &pb.Data{}
		data := buff[:n]
		err = proto.Unmarshal(data, rec)
		if err != nil {
			panic(err)
		}
		fmt.Println("接收到数据：", conn.RemoteAddr(), rec)

		//send, err := proto.Marshal(rec)
		//if err != nil {
		//	panic(err)
		//}

		//fmt.Println(send)
		//conn.Write(send)
		//fmt.Println("Server send ovwr")
	}
}

func handleConn(conn net.Conn) {
	clients[conn.RemoteAddr().String()] = conn
}
