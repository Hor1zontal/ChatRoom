package utils

import (
	common "common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Dispatcher struct {
	Conn net.Conn
	Buf [10240]byte
}

func (dispatcher Dispatcher) ReadData() (message common.Message, err error) {
	buf := dispatcher.Buf

	//读取消息长度信息
	n, err := dispatcher.Conn.Read(buf[:4])
	if err != nil {
		return
	}
	var dataLen uint32
	dataLen = binary.BigEndian.Uint32(buf[0:4])

	//读取消息本体
	if err != nil {
		fmt.Printf("server read data lgoin data error : %v", err)
	}

	//对比消息本身的长度和期望长度是否匹配
	if n != int(dataLen) {
		err = errors.New("login message length error")
		return
	}

	//从conn中解析消息并存放到message中，此处一定传递的是message的地址
	err = json.Unmarshal(buf[:dataLen], &message)
	if err != nil {
		fmt.Printf("json.Unmarshl error: %v", err)
	}
	return
}
func (dispatcher Dispatcher) WriteData(data []byte) (err error) {
	var dataLen uint32
	dataLen = uint32(len(data))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4], dataLen)

	// 将消息长度发送给客户端
	_, err = dispatcher.Conn.Write(bytes[:])
	if err != nil {
		fmt.Printf("send data length to server error: %v\n", err)
		return
	}

	// 发送消息本身给客户端
	_, err = dispatcher.Conn.Write(data)
	if err != nil {
		fmt.Printf("send data length to server error: %v", err)
		return
	}
	return
}
