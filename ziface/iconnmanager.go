package ziface

type IConnManager interface {

    Add(c IConnection)

    Remove(c IConnection)

    Get(cid uint32) (IConnection, error)

    Len() int

    ClearConn()

}