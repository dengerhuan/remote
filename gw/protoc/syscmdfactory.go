package protoc

import EventBus "gw/eventbus"

var SysCmdFactory = sysCmdFactory{}

var eventbus = EventBus.GlobalBus

type sysCmdFactory struct {
}

func (u sysCmdFactory) GetCmd(cmd uint16) Command {

	//return
	return DefaultCommand{}
}
