package znet

import (
    "errors"
    "fmt"
    "zinx/ziface"
)

type MsgHandler struct {
    Apis map[uint32]ziface.IRouter
    WorkerPoolSize uint32                    //业务工作Worker池的数量
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

func newMsgHandler() *MsgHandler {
    return &MsgHandler{
        Apis:make(map[uint32]ziface.IRouter),
    }
}

