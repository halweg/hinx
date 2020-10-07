package znet

import (
	"fmt"
	"net"
	"zinx/izface"
)

type Connection struct {

	Connection *net.TCPConn

	ConnectionID uint32

	isClose bool

	ExitChan chan bool

	Router izface.IRouter

}

func (c *Connection) StartReader() {
	fmt.Println("reader 读业务启动中.....")
	defer fmt.Println("connID=", c.ConnectionID, "Reader IS " +
		"exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Connection.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}

		/*if err := c.HandleAPI(c.Connection, buf, cnt); err != nil {
			fmt.Println("connid ", c.ConnectionID, "handler error", err)
			break
		}*/

		req := Request{
			Conn: c,
			data: buf,
		}

		go func(request izface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)


	}
}

func (c *Connection) Start() {

	fmt.Println("Connection creating...... connctionID = ", c.ConnectionID)

	go c.StartReader()


	//TODO 启动从当前连接写数据的业务
}


func (c *Connection) Stop() {
	fmt.Println("coonID 为 %d 的连接正在关闭中......", c.ConnectionID)
	if c.isClose == true {
		return
	}

	c.isClose = true

	c.Connection.Close()

	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Connection
}

func (c *Connection) RemoteAddr() net.Addr{
	return c.Connection.RemoteAddr()
}


func (c *Connection) GetConnectID() uint32 {
	return c.ConnectionID
}

func (c *Connection) Send([]byte) error {
	return nil
}

func NewConnection(connection *net.TCPConn, connectionID uint32, router izface.IRouter) *Connection {

	c := &Connection{
		Connection:   connection,
		ConnectionID: connectionID,
		isClose:      false,
		ExitChan:     make(chan bool, 1),
		Router: router,
	}

	return c
}