package log

import "fmt"

func Log(a ...interface{}) {
	fmt.Println("[Server]:", a)
}

func Loge(a ...interface{}) {
	fmt.Sprint(a)
}
