package main

import (
    "fmt"
    "io"
    "math/rand"
    "net"
    "time"
    "zinx/znet"
)



func main() {

	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")

	if err != nil {
		panic(err)
	}

	for {

	    dp := znet.NewDadaPack()
        seed := time.Now().UnixNano()
        r1 := rand.New(rand.NewSource(seed))
        msgId := uint32(r1.Intn(2))
        fmt.Println(msgId)
        msg , er := dp.Pack(znet.NewMsgPackage(msgId, []byte("连通性测试....")))
	    if er != nil {
	        panic(err)
        }
		_, err := conn.Write(msg)

		if err != nil {
			panic(err)
		}

		respHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, respHead); err != nil {
		    panic(err)
        }

        respMessage, err := dp.UnPack(respHead)
        if err != nil {
            panic(err)
        }
        if respMessage.GetMessageLen() > 0 {
            msg := respMessage.(*znet.Message)
            msg.Data = make([]byte, msg.GetMessageLen())

            _, err := io.ReadFull(conn, msg.Data)

            if err != nil {
                panic(err)
            }

            fmt.Printf("接收到服务器端的返回: %s, 消息长度为 %d " +
                "消息ID为 %d \n", string(msg.Data), msg.DataLen, msg.Id)

        }


		time.Sleep(time.Second)

	}




}
