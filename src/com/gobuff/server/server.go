package server

import (
	"net"
	"gobuff/src/com/gobuff/config"
	"strconv"
	"gobuff/src/com/gobuff/log"
	"gobuff/src/com/gobuff/connection"
	"gobuff/src/com/gobuff/cache"
)

func Run(con config.Config) {
	addr := con.ServerIp + ":" + strconv.Itoa(con.Port)
	listener, err := net.Listen(con.Protocol, addr)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	log.Log("服务器启动成功，等待连接")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		log.Log("新连接：", conn.RemoteAddr())

		cache.Clients[conn.RemoteAddr().String()] = conn
		go connection.NewConnection(conn)
	}
}
