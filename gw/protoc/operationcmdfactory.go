package protoc

import "log"

var OperationCmdFactory = operationCmdFactory{}

type operationCmdFactory struct {
}

func (u operationCmdFactory) GetCmd(cmd uint16) Command {

	log.Println(cmd)
	switch cmd {

	case 1:
		return srdCheck{}
	case 3:
		return srdStart{}

	case 4, 5:
		return erdApplyReq{}
	case 7:
		return erdStop{}
	case 9:
		return erdAckReq{}
	}

	return srdCheck{}
}
