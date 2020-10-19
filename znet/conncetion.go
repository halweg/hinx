package znet

import (
    "errors"
    "fmt"
    "io"
    "net"
    "sync"
    "zinx/utils"
    "zinx/ziface"
)

type Connection struct {
    TCPServer ziface.IServer

    Connection *net.TCPConn

    ConnectionID uint32

    isClose bool

    ExitChan chan bool

    msgChan chan []byte

    MsgHandler  ziface.IMsgHandler

    //用户自定义连接属性
    Property map[string]interface{}

    PropertyLock sync.RWMutex
}

func (c *Connection) StartReader() {
	fmt.Println("Conn  StartReader() is running.....")
	defer fmt.Println("connID = " , c.ConnectionID, "StartReader() IS " +
		"exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {

		dp := NewDadaPack()

		headData := make([]byte, dp.GetHeadLen())
		if _, err :=  io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head err\n", err)
            return
		}

		msg,  err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack err \n", err)
            return
		}

		var data []byte
		if msg.GetMessageLen() > 0 {
			data = make([]byte, msg.GetMessageLen())

			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data err", err)
                return
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

		if utils.GlobalObject.WorkerPoolSize > 0 {
		    c.MsgHandler.SendMsgToTasKQueue(&req)
        } else {
            go c.MsgHandler.DoMessageHandler(&req)
        }
	}
}

func (c *Connection) StartWrite() {

    fmt.Println("Conn StartWrite() :" , c.Connection.RemoteAddr().String())
    defer fmt.Println("Conn StartWrite() close : ", c.Connection.RemoteAddr().String())

    for {
        select {
        case data := <-c.msgChan :
            if _, err := c.Connection.Write(data); err != nil {
                fmt.Println("Conn StartWrite() err :", err)
                return
            }
        case <-c.ExitChan :
            return
        }
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

	c.msgChan <- binMsg

	return nil

}

func (c *Connection) Start() {

	fmt.Println("Conn Start()...... connID = ", c.ConnectionID)

	go c.StartReader()

    go c.StartWrite()

	//调用hook
    c.TCPServer.CallOnnStart(c)
}


func (c *Connection) Stop() {
	fmt.Println("Conn Stop()...... connID =", c.ConnectionID)
	if c.isClose == true {
		return
	}

	c.isClose = true

	//调用hook
	c.TCPServer.CallOnConnStop(c)

	c.Connection.Close()
    c.ExitChan <- true

    c.TCPServer.GetConnManager().Remove(c)

	close(c.ExitChan)
	close(c.msgChan)
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

func (c *Connection) SetProperty(key string, v interface{}) {
    c.PropertyLock.Lock()
    defer c.PropertyLock.Unlock()
    c.Property[key] = v
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
    c.PropertyLock.RLock()
    defer c.PropertyLock.RUnlock()

    if v , e := c.Property[key]; e {
        return v, nil
    }

    fmt.Println("property ", key, "not found!")
    return nil, errors.New("property "+key+" not found!")

}

func (c *Connection) RemoveProperty(key string) {
    c.PropertyLock.Lock()
    defer c.PropertyLock.Unlock()

    delete(c.Property, key)
}


func NewConnection(s ziface.IServer,connection *net.TCPConn, connectionID uint32,  handler ziface.IMsgHandler) *Connection {

	c := &Connection{
	    TCPServer:s,
		Connection:   connection,
		ConnectionID: connectionID,
		isClose:      false,
		msgChan:      make(chan []byte),
		ExitChan:     make(chan bool, 1),
		MsgHandler:   handler,
		Property:     make(map[string]interface{}),
	}

	c.TCPServer.GetConnManager().Add(c)

	return c
}