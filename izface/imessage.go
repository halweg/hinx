package izface

/*
*/
type Imessage interface {
	GetMessageId() uint32

	GetMessageLen() uint32

	GetData() []byte

	SetMessageData([]byte)

	SetDataLen(uint32)
}