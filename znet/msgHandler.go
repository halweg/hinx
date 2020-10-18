package znet

import (
    "errors"
    "fmt"
    "zinx/utils"
    "zinx/ziface"
)

type MsgHandler struct {
    Apis map[uint32]ziface.IRouter
    WorkerPoolSize uint32                    //业务Worker的数量
    TaskQueue      []chan ziface.IRequest    //Worker负责取任务的消息队列
}

func (m *MsgHandler) DoMessageHandler(req ziface.IRequest) {

    if _, ok := m.Apis[req.GetMsgID()]; !ok{
        fmt.Printf("messageID %d router not found \n", req.GetMsgID())
        req.GetConnection().SendMsg(404, []byte("Route not found"))
        return
    }

    m.Apis[req.GetMsgID()].PreHandle(req)
    m.Apis[req.GetMsgID()].Handle(req)
    m.Apis[req.GetMsgID()].PostHandle(req)

}

func (m *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter)  {

    if _, ok := m.Apis[msgID]; ok {
        err := errors.New( fmt.Sprintf("the router of msgID %d has exits!", msgID))
        panic(err)
    }

    m.Apis[msgID] = router

}

func (m *MsgHandler) StartWorkerPool() {

    for i :=0; i < int(utils.GlobalObject.WorkerPoolSize); i ++ {
        m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerPoolSize)
        go m.StartOneWorker(i, m.TaskQueue[i])
    }

}

func (m *MsgHandler) StartOneWorker(workerID int, queue chan ziface.IRequest) {
    fmt.Println("MsgHandler StartOneWorker() running...,  WorkerID is ", workerID)
    for {
        select {
        case re := <-queue:
            m.DoMessageHandler(re)
        }
    }
}

func (m *MsgHandler) SendMsgToTasKQueue(re ziface.IRequest) {
    workerID := re.GetConnection().GetConnectID() % utils.GlobalObject.WorkerPoolSize
    fmt.Println("connID : ", re.GetConnection().GetConnectID(), " SendMsgToTasKQueue() to TaskQueue, workerID = ", workerID)
    m.TaskQueue[workerID] <- re
}


func newMsgHandler() *MsgHandler {
    return &MsgHandler{
        Apis:make(map[uint32]ziface.IRouter),
        WorkerPoolSize : utils.GlobalObject.WorkerPoolSize,
        TaskQueue: make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
    }
}


