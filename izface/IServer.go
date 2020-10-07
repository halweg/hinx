package izface

type IServer interface {
	Start()

	Stop()

	Server()

	AddRouter(router IRouter)
}

