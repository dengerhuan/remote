package protoc

import (
	"encoding/binary"
	"fmt"
	"github.com/go-netty/go-netty"
)

/**
codec 0 json
      1 bin
      2 string
*/

const (
	TypeIndex     = 6
	DomainIndex   = 7
	CmdCodeIndex = 8 // 16
	SwitchIndex   = 10
	CodecIndex    = 11
	LenIndex      = 12 //32
	TxId          = 16
)

// cmd interface
type Command interface {
	Execute(ctx netty.InboundContext, message netty.Message)
}

type DefaultCommand struct {
}

func (d DefaultCommand) Execute(ctx netty.InboundContext, message netty.Message) {
	fmt.Println("default cmd")

}

// cmd factory interface
type CmdFactory interface {
	GetCmd(cmd uint16) Command
}

func DecodeHead(msg []byte) (uint8, uint32, uint8) {
	flag := msg[11]
	l := binary.BigEndian.Uint32(msg[12:16])
	codec := msg[11]
	return flag, l, codec
}
