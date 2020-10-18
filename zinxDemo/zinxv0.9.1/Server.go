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

func HookBeforeConn(conn ziface.IConnection) {
    fmt.Println("Hook func on start is running....")
    if err := conn.SendMsg(202, []byte("hello welcome call HookFunc on conn start\n")); err!= nil{
        fmt.Println(err)
    }
}

func HookAfterConn(conn ziface.IConnection) {
    fmt.Println("Hook func on stop func is running....")
}

func main() {

	s := znet.NewZinxServer("[zinx0.9.1]")

	s.SetOnConnStart(HookBeforeConn)
	s.SetOnConnStop(HookAfterConn)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.Server()

}
