package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
	"zinx/utils"
)

type Server struct {

	Name string

	IP string

	IPVersion string

	port string

    MsgHandler ziface.IMsgHandler

	ConnManager ziface.IConnManager

}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
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
	fmt.Printf("[Zinx] server name is : %s, listenner at Ip:%s:%s",
		utils.GlobalObject.Name,
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort)

	fmt.Printf("[Zinx] Version is %s, MaxPackageSize is %d, MaxConnNum is %d",
		utils.GlobalObject.Version, utils.GlobalObject.MaxPackageSize,
		utils.GlobalObject.MaxConn)

	go func() {

	    s.MsgHandler.StartWorkerPool()
	    
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


            if s.ConnManager.Len() >= utils.GlobalObject.MaxConn {
                fmt.Println("===========================to many conn.....===============================")
                conn.Close()
                continue
            }


            dealConnection := NewConnection(s ,conn, cid, s.MsgHandler)
			cid ++
            go dealConnection.Start()
		}

	}()




}

func (s *Server) Stop() {
    fmt.Println("server is stopping...")
    s.ConnManager.ClearConn()
}

func (s *Server) Server() {

	fmt.Println("服务启动中.......")
	s.Start()
	fmt.Println("服务启动成功！")
	select {}
}


func (s *Server) GetConnManager () ziface.IConnManager {
    return s.ConnManager
}

func NewZinxServer(name string) ziface.IServer {
	return &Server{
		Name: utils.GlobalObject.Name,
		IP: utils.GlobalObject.Host,
		IPVersion: "tcp4",
		port: utils.GlobalObject.TcpPort,
		MsgHandler: newMsgHandler(),
        ConnManager: NewConnManager(),
	}
}