package main

import (
	"gobuff/src/com/gobuff/server"
	"gobuff/src/com/gobuff/config"
)

func main() {
	con := config.Config{Protocol: "tcp", Port: 8545, ServerIp: "127.0.0.1"}
	server.Run(con)
}
