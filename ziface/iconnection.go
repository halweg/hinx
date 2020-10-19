package ziface

import "net"

type IConnection interface {
	Start()
	
	Stop()

	GetTCPConnection() *net.TCPConn

	GetConnectID() uint32

	Send([]byte) error

	RemoteAddr() net.Addr

	SendMsg(msgID uint32, data []byte) error

    SetProperty(key string, v interface{})

	GetProperty(key string) (interface{}, error)

    RemoveProperty(key string)
}

type HandlerFunc func(*net.TCPConn, []byte, int) error