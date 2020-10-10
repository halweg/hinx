package ziface

type IServer interface {
	Start()

	Stop()

	Server()

	AddRouter(msgID uint32, router IRouter)
}

