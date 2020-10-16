package ziface

type IMsgHandler interface {

    DoMessageHandler(request IRequest)

    AddRouter(msgID uint32, router IRouter)

    StartWorkerPool()

    StartOneWorker(workerID int, queue chan IRequest)

    SendMsgToTasKQueue(re IRequest)
}

