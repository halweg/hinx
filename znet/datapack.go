package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/ziface"
	"zinx/utils"
)

//封包和拆包的模块

type DataPack struct {}


func NewDadaPack() *DataPack {
	return  &DataPack{}
}

func (dp *DataPack) GetHeadLen () uint32 {
	//魔法数字，
	//msg 头部 4字节表示长度，4字节表示消息ID
	return 8
}


func (dp *DataPack) Pack(msg ziface.Imessage) ([]byte, error) {

	databuf := bytes.NewBuffer([]byte{})

	err := binary.Write(databuf, binary.LittleEndian, msg.GetMessageLen())

	if err != nil {
		return nil, err
	}

	err = binary.Write(databuf, binary.LittleEndian, msg.GetMessageId())

	if err != nil {
		return nil, err
	}

	err = binary.Write(databuf, binary.LittleEndian, msg.GetData())

	if err != nil {
		return nil, err
	}

	return databuf.Bytes(), nil

}

func (dp *DataPack) UnPack(binData []byte) (ziface.Imessage, error) {

	dataBuf := bytes.NewReader(binData)

	msg := &Message{}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	if utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg")
	}

	return msg, nil

}