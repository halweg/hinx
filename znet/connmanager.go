package znet

import (
    "errors"
    "fmt"
    "sync"
    "zinx/ziface"
)

type ConnManager struct {
    connection map[uint32] ziface.IConnection
    connLock sync.RWMutex
}

func (c2 *ConnManager) Add(c ziface.IConnection) {
    c2.connLock.Lock()
    defer c2.connLock.Unlock()
    c2.connection[c.GetConnectID()] = c
    fmt.Println("conn add to connManger success, conn num is ", c.GetConnectID())
}

func (c2 *ConnManager) Remove(c ziface.IConnection) {
    c2.connLock.Lock()
    defer c2.connLock.Unlock()
    delete(c2.connection, c.GetConnectID())
    fmt.Println("conn ", c.GetConnectID(), "remove success form connManager")
}

func (c2 *ConnManager) Get(cid uint32) (ziface.IConnection, error) {
    c2.connLock.RLock()
    defer c2.connLock.RUnlock()
    if conn, ok := c2.connection[cid]; ok {
        return conn, nil
    } else {
        return nil, errors.New("connection NOTFOUND")
    }
}

func (c2 *ConnManager) Len() int {
   return len(c2.connection)
}

func (c2 *ConnManager) ClearConn() {
    c2.connLock.Lock()
    defer c2.connLock.Unlock()
    for k, v:= range c2.connection {
        v.Stop()
        delete(c2.connection, k)
    }
    fmt.Println("all connections clear from connManager")
}

func NewConnManager () ziface.IConnManager {
    return &ConnManager{
        connection: make(map[uint32] ziface.IConnection),
    }
}
