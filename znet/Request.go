package znet

import "zinx/izface"

type Request struct {

	Conn izface.IConnection

	data []byte

}

func (r *Request) GetConnection() izface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.data
}