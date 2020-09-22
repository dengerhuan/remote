package protoc

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
