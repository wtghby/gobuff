package main

import (
	"time"
	"fmt"
)

func main() {
	hh()
}

func hh() {
	tick := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-tick.C:
			fmt.Println("执行")
		}
	}
}
