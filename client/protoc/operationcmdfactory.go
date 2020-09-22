package protoc

import (
	"client/drivemanager"
	"github.com/go-netty/go-netty"
	"log"
)

var OperationCmdFactory = operationCmdFactory{}

type operationCmdFactory struct {
}

func (u operationCmdFactory) GetCmd(cmd uint16) Command {

	switch cmd {
	case 0:
		return srdCheck{}
	case 2:
		return srdStart{}
	case 4:
		return erdApply{}
	case 6:
		return erdStop{}
	case 8:
		return erdAck{}
	}
	return DefaultCommand{}
}

type erdApply struct {
}

func (d erdApply) Execute(ctx netty.InboundContext, message netty.Message) {
	log.Println("stop")
	//log.Println(message)
	drivemanager.RdState = false
}

type erdStop struct {
}

func (d erdStop) Execute(ctx netty.InboundContext, message netty.Message) {

}

type erdAck struct {
}

func (d erdAck) Execute(ctx netty.InboundContext, message netty.Message) {

}
