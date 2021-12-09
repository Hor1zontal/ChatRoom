package process

import (
	common "common/message"
	"fmt"
	"io"
	"net"
	"server/model"
	"server/utils"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) messageProcess(message common.Message) (err error){
	switch message.Type {
	case common.LoginMessageType:
	case common.RegisterMessageType:
	case common.LoginResponseMessageType:
	case common.RegisterResponseMessageType:

	default:
		fmt.Printf("other type\n")
	}
	return
}

func (this *Processor) MainProcess() {
	for {

		dispatcher := utils.Dispatcher{Conn: this.Conn}
		message, err := dispatcher.ReadData()
		if err != nil {
			if err == io.EOF {
				cc := model.ClientConn{}
				cc.Del(this.Conn)
				fmt.Printf("client closed!\n")
				break
			}
			fmt.Printf("get login message error: %v", err)
		}
		err = this.messageProcess(message)
		if err != nil {
			fmt.Printf("some error: %v\n", err)
			break
		}
	}
}