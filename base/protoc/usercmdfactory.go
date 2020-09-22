package protoc

import (
	"base/authz"
	"base/session"
	"encoding/binary"
	"github.com/go-netty/go-netty"
	"log"
)

// domain user=2
// type =cmd=0
//
var hi [30]byte
var UserCmdFactory = userCmdFactory{}

type userCmdFactory struct {
}

func (u userCmdFactory) GetCmd(cmd uint16) Command {

	switch cmd {
	case 0:
		return UserHi{}
	case 2:
		return UserRegister{}
	}

	return DefaultCommand{}
}

// hi cmdcodec=0
// hi cmdres=1



type UserHi struct {
}

func (d UserHi) Execute(ctx netty.InboundContext, message netty.Message) {
	//fmt.Println("user hi")

	//  0 0 0 0     - 0  0     0 2 0 1 0 1 0 0 0 1 0 0 0 0 0]

	// handler hi
	msg := message.([]byte)
	_, l, codec := DecodeHead(msg)

	if codec == 2 {
		s := string(msg[20 : 20+l])

		_session := session.SessionManagerInst.SessionAt(ctx.Channel().ID())

		res := hi[:20]

		res[7] = 2
		res[CmdCodeIndex+1] = 1
		res[CodecIndex] = 1
		res[LenIndex+3] = 1

		_, ok := _session.Attr(s)

		if ok {
			ctx.Write(append(res, 1))
		} else {
			ctx.Write(append(res, 0))
		}

	}
}

// hi cmdcodec=0
// hi cmdres=1
type UserRegister struct {
}

func (d UserRegister) Execute(ctx netty.InboundContext, message netty.Message) {

	//fmt.Println("user register")
	msg := message.([]byte)
	_, l, codec := DecodeHead(msg)

	if codec == 2 {
		s := string(msg[20 : 20+l])

		vid, ok := authz.Auth(ctx.Channel().ID(), s)

		res := hi[:21]

		res[CmdCodeIndex+1] = 3
		res[CodecIndex] = 1
		res[LenIndex+3] = 1

		if ok {

			log.Println("device", vid.GetId(), "register  success")

			res[20] = 1

			tmpid := []byte(vid.GetId())

			//fmt.Println(tmpid)

			binary.BigEndian.PutUint32(res[LenIndex:LenIndex+4], uint32(1+len(tmpid)))

			ctx.Write(append(res, tmpid...))
		} else {
			res[20] = 0
			ctx.Write(res)
		}
	}
}
