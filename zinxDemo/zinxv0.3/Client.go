package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")

	if err != nil {
		panic(err)
	}

	for {

		_, err := conn.Write([]byte("你好，zinx"))

		if err != nil {
			panic(err)
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)

		if err != nil {
			panic(err)
		}

		fmt.Printf("接收到服务器端的返回: %s\n", buf[:cnt])

		time.Sleep(time.Second)

	}




}
