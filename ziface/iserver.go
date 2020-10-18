package ziface

type IServer interface {
	Start()

	Stop()

	Server()

	AddRouter(msgID uint32, router IRouter)

    GetConnManager () IConnManager

    SetOnConnStart (func(conn IConnection))

    SetOnConnStop (func(conn IConnection))

    CallOnnStart (conn IConnection)

    CallOnConnStop (conn IConnection)

}

