package znet

import (
	"fmt"
	"net"
	"zinx/izface"
)

type Server struct {

	Name string

	IP string

	IPVersion string

	port string

	Router izface.IRouter

}

func (s *Server) AddRouter(router izface.IRouter) {

	s.Router = router
	fmt.Println("add Router success")
}

/*func CallBackToClient(conn *net.TCPConn, data []byte, cnt int)  error {
	fmt.Println("[Conn handle] CallbackToClient...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("Write back buf err", err)
		return errors.New("CallBackTOClient error")
	}

	return nil
}*/

func (s *Server) Start() {
	fmt.Printf("start server : %s at %s:%s\n", s.Name, s.IP, s.port)

	go func() {

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%s", s.IP, s.port))

		if err != nil {
			fmt.Printf("服务启动失败！！")
			panic(err)
		}

		listener, err := net.ListenTCP("tcp",addr)

		if err != nil {
			panic(err)
		}
		var cid uint32
		cid = 0
		for  {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println(err)
				continue
			}
			dealConnection := NewConnection(conn, cid, s.Router)
			cid ++
			go dealConnection.Start()
		}

	}()




}

func (s *Server) Stop() {

}

func (s *Server) Server() {

	fmt.Println("服务启动中.......")
	s.Start()
	fmt.Println("服务启动成功！")
	select {}
}

func NewZinxServer(name string) izface.IServer {
	return &Server{
		Name: name,
		IP:        "127.0.0.1",
		IPVersion: "tcp4",
		port:      "8999",
		Router: nil,
	}
}