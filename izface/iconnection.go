package izface

import "net"

type IConnection interface {
	Start()
	
	Stop()

	GetTCPConnection() *net.TCPConn

	GetConnectID() uint32

	Send([]byte) error

	RemoteAddr() net.Addr

}

type HandlerFunc func(*net.TCPConn, []byte, int) error