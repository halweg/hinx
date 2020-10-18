package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T)  {


	listener, err := net.Listen("tcp", "127.0.0.1:8999")

	if err != nil {
		fmt.Println("server listen err :", err)
		return
	}

	go func() {
		for  {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept err", err)
				return
			}

			go func(conn net.Conn) {

				//读2次，
				dp := NewDadaPack()

				for  {

					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head err", err)
						break
					}

					msgHead, err := dp.UnPack(headData)
					if err != nil {
						fmt.Println("server unpack err", err)
						return
					}

					if msgHead.GetMessageLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetMessageLen())

						_, err := io.ReadFull(conn, msg.Data)
						if err != nil {
							fmt.Println("server unpack err")
							return
						}

						fmt.Println("---->Recv MsgId:", msg.Id, ",dataLen = ", msg.DataLen,
							",data=", string(msg.Data))


					}

				}

			}(conn)
		}
	}()


	connClient, err := net.Dial("tcp", "127.0.0.1:8999")

	if err != nil {
		fmt.Println("client dial error ", err)
		return
	}

	dp := NewDadaPack()

	msg1 := &Message{
		Id : 1,
		DataLen: 4,
		Data : []byte{'z', 'i', 'n', 'x'},
	}

	sendData1 , err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
	}

	msg2 := &Message{
		Id : 1,
		DataLen: 5,
		Data : []byte{'h', 'e', 'l', 'l', '0'},
	}

	sendData2 , err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg1 error", err)
	}

	sendData1 = append(sendData1, sendData2...)

	connClient.Write(sendData1)

	select {}

}
