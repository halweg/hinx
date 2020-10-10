package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(request ziface.IRequest ) {
	fmt.Println("call router PreHandle......")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("调用了PreHandle\n"))
	if err != nil {
		fmt.Println("call back PreHandle error")
	}

}

func (pr *PingRouter) Handle(request ziface.IRequest ) {
	fmt.Println("call router handle......")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("调用了handle\n"))
	if err != nil {
		fmt.Println("call back handle error")
	}
}

func (pr *PingRouter) PostHandle(request ziface.IRequest ) {
	fmt.Println("call router PostHandle......")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("调用了PostHandle\n"))
	if err != nil {
		fmt.Println("call back PostHandle error")
	}
}

func main() {

	s := znet.NewZinxServer("[zinx0.3]")

	s.AddRouter(&PingRouter{})

	s.Server()

}
