package ziface

type IMsgHandler interface {

    DoMessageHandler(request IRequest)

    AddRouter(msgID uint32, router IRouter)

}

