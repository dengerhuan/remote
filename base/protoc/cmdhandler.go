package protoc

import (
	"encoding/binary"
	"fmt"
	"github.com/go-netty/go-netty"
)

type CmdHandler struct {
}

func (CmdHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {

	// check if this is a  cmd
	//fmt.Println("cmd handler read")

	msg := message.([]byte)
	if len(msg) < 16 {
		return
	}
	msgType := msg[6]

	switch msgType {
	case 0:
		cmdFactory := checkDomain(msg[7])

		if cmdFactory != nil {
			cmdFactory.GetCmd(binary.BigEndian.Uint16(msg[8:10])).Execute(ctx, msg)
		} else {
			ctx.HandleRead(message)
		}

	case 1:
		ctx.HandleRead(msg)
	default:
		fmt.Println("unsupport ")
	}

}

func checkDomain(domain uint8) CmdFactory {

	switch domain {
	case 0:
		return nil
	case 1:
		return NtpCmdFactory
	case 2:
		return UserCmdFactory
	case 3:
		return OperationCmdFactory

	default:
		fmt.Println("unSupport domain")
		return nil
	}
}
