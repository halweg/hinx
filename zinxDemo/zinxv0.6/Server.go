package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) Handle(request ziface.IRequest ) {
	fmt.Println("call router handle......")
	fmt.Println(" Hi WelCome use MsgId ", request.GetMsgID(), " ，you msg \"  ", string(request.GetData()), "\" is success!")

	request.GetConnection().SendMsg(1, []byte("ping ping ping"))
}

type HelloRouter struct {
    znet.BaseRouter
}

func (hr *HelloRouter) Handle(request ziface.IRequest) {
    request.GetConnection().SendMsg(1, []byte("Hello"))
}

func main() {

	s := znet.NewZinxServer("[zinx0.6]")

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Server()

}
