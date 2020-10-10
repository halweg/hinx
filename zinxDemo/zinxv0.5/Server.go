package main

import (
	"fmt"
	"zinx/izface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handle(request izface.IRequest ) {
	fmt.Println("call router handle......")
	fmt.Println(" Hi WelCome use MsgId ", request.GetMsgID(), " ，you msg \"  ", string(request.GetData()), "\" is success!")

	request.GetConnection().SendMsg(1, []byte("你好，欢迎光临"))
}

func main() {

	s := znet.NewZinxServer("[zinx0.5]")

	s.AddRouter(&PingRouter{})

	s.Server()

}
