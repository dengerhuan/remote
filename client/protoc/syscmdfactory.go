package protoc

var SysCmdFactory = sysCmdFactory{}

type sysCmdFactory struct {
}

func (u sysCmdFactory) GetCmd(cmd uint16) Command {
	//return
	return DefaultCommand{}
}



