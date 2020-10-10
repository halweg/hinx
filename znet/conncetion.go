package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

type Connection struct {

	Connection *net.TCPConn

	ConnectionID uint32

	isClose bool

	ExitChan chan bool

	//Router ziface.IRouter

	MsgHandler ziface.IMsgHandler
}

func (c *Connection) StartReader() {
	fmt.Println("reader 读业务启动中.....")
	defer fmt.Println("connID=", c.ConnectionID, "Reader IS " +
		"exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		/*buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Connection.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			continue
		}*/

		/*if err := c.HandleAPI(c.Connection, buf, cnt); err != nil {
			fmt.Println("connid ", c.ConnectionID, "handler error", err)
			break
		}*/

		dp := NewDadaPack()

		headData := make([]byte, dp.GetHeadLen())
		if _, err :=  io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head err\n", err)
			break
		}

		msg,  err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack err \n", err)
			break
		}

		var data []byte
		if msg.GetMessageLen() > 0 {
			data = make([]byte, msg.GetMessageLen())

			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data err", err)
				break
			}
		}
		msg.SetMessageData(data)

		//读取客户端的 msg 头 二进制流的前8个字节

		//拆包，得到 msgId 和 msgDataLen, 放到 msg 消息中

		//根据 dataLen , 再次读取Data, 放在 msgData中

		req := Request{
			Conn: c,
			msg: msg,
		}
		
		go c.MsgHandler.DoMessageHandler(&req)


	}
}

//把要发给客户端的消息先封包再发送
func (c *Connection) SendMsg(messageID uint32, data []byte) error {

	if c.isClose == true {
		return errors.New("connection is close when send msg")
	}

	dp := NewDadaPack()

	binMsg, err := dp.Pack(NewMsgPackage(messageID, data))
	if err != nil {
		fmt.Println("pack err msg id =", messageID)
		return errors.New("pack error msg")
	}

	if _, err := c.Connection.Write(binMsg); err != nil {
		fmt.Println("write msg id ", messageID, " is error")
		return errors.New("conn write msg error")
	}

	return nil

}

func (c *Connection) Start() {

	fmt.Println("Connection creating...... connctionID = ", c.ConnectionID)

	go c.StartReader()


	//TODO 启动从当前连接写数据的业务
}


func (c *Connection) Stop() {
	fmt.Printf("coonID 为 %d 的连接正在关闭中......", c.ConnectionID)
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

func NewConnection(connection *net.TCPConn, connectionID uint32,  handler ziface.IMsgHandler) *Connection {

	c := &Connection{
		Connection:   connection,
		ConnectionID: connectionID,
		isClose:      false,
		ExitChan:     make(chan bool, 1),
		MsgHandler:   handler,
	}

	return c
}