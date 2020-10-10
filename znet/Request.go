package znet

import "zinx/ziface"

type Request struct {

	Conn ziface.IConnection

	msg ziface.Imessage

}

func (r *Request) GetConnection() ziface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMessageId()
}