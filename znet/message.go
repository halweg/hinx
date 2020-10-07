package znet

type Message struct {

	Id uint32

	DataLen uint32

	Data []byte
}

func NewMsgPackage(id uint32, data []byte)*Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMessageId() uint32 {
	return m.Id
}

func (m *Message) GetMessageLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMessageData(b []byte) {
	m.Data = b
}

func (m *Message) SetDataLen(l uint32) {
	m.DataLen = l
}
