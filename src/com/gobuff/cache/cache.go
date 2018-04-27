package cache

import "net"

var Clients = make(map[string]net.Conn)
