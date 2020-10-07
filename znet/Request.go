package znet

import "zinx/izface"

type Request struct {

	Conn izface.IConnection

	msg izface.Imessage

}

func (r *Request) GetConnection() izface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMessageId()
}