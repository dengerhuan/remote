package protoc

import (
	"encoding/binary"
	"github.com/go-netty/go-netty"
	"time"
)

// domain user=2
// type =cmd=0
//
var ntpBuf [52]byte
var NtpCmdFactory = ntpCmdFactory{}

type ntpCmdFactory struct {
}

func (u ntpCmdFactory) GetCmd(cmd uint16) Command {

	switch cmd {
	case 0:
		return NetTimeStamp{}
	}

	return DefaultCommand{}
}

// hi cmdcodec=0
// hi cmdres=1
type NetTimeStamp struct {
}

func (d NetTimeStamp) Execute(ctx netty.InboundContext, message netty.Message) {

	//fmt.Println(message)
	ntp := ntpBuf[:]

	// handler hi
	msg := message.([]byte)
	_, _, codec := DecodeHead(msg)

	if codec == 1 {

		ntp[TypeIndex] = 0
		ntp[DomainIndex] = 1
		ntp[CodecIndex] = 1
		ntp[CmdCodeIndex+1] = 1

		r := binary.BigEndian.Uint64(msg[20:28])
		binary.BigEndian.PutUint32(ntp[LenIndex:LenIndex+4], 32)
		binary.BigEndian.PutUint64(ntp[20:28], r)
		binary.BigEndian.PutUint64(ntp[28:36], uint64(time.Now().UnixNano()))
		binary.BigEndian.PutUint64(ntp[36:44], uint64(time.Now().UnixNano()))
		ctx.Write(ntp)
	}
}
