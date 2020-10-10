package ziface

type IDataPack interface {
	GetHeadLen() uint32

	Pack(msg Imessage) ([]byte, error)

	UnPack([]byte) (Imessage, error)
}