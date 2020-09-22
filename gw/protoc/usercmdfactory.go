package protoc

import (
	"encoding/binary"
	"fmt"
	"github.com/go-netty/go-netty"
	"gw/authz"
)

// domain user=2
// type =cmd=0
//
var UserCmdFactory = userCmdFactory{}

type userCmdFactory struct {
}

func (u userCmdFactory) GetCmd(cmd uint16) Command {

	switch cmd {
	case 1:
		return userHiResponse{}
	case 3:
		return userRegisterResponse{}
	}
	return DefaultCommand{}
}

// hi cmdcodec=0
// hi cmdres=1
type userHiResponse struct {
}

func (userHiResponse) Execute(ctx netty.InboundContext, message netty.Message) {

	msg := message.([]byte)
	_, _, codec := DecodeHead(msg)

	if codec == 1 {

		if msg[20] == 0 {
			fmt.Println("server response  not register user")
			authz.Register(ctx.Channel())
		}
	}
}

type userRegisterResponse struct {
}

func (userRegisterResponse) Execute(ctx netty.InboundContext, message netty.Message) {

	msg := message.([]byte)
	_, l, codec := DecodeHead(msg)

	if codec == 1 {
		if msg[20] == 1 {
			//fmt.Println(msg)
			tmpId := string(msg[21 : 20+l])
			fmt.Println("register successful device temp id:", tmpId)

			ctx.Write(reportCategory(msg[21 : 20+l]))

		}
		if msg[20] == 0 {
			fmt.Println("register error")
		}
	}
}

var category [28]byte


//logicType     uint8 // 21 类型  --console2 --车0 - 台驾1  3 monitor
//category      uint8 //20  能力
func reportCategory(tepid []byte) []byte {

	msg := category[:22]
	msg[TypeIndex] = 1
	msg[DomainIndex] = 3
	msg[CmdCodeIndex+1] = 0
	msg[CodecIndex] = 1

	binary.BigEndian.PutUint32(msg[LenIndex:LenIndex+4], uint32(len(tepid)))

	msg[20] = 0

	msg[21] = 2
	return append(msg, tepid...)

}
