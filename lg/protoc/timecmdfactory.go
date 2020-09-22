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

/**
  Originate Timestamp       T1        客户端发送请求的时间
  Receive Timestamp        T2        服务器接收请求的时间
  Transmit Timestamp       T3        服务器答复时间
  Destination Timestamp     T4        客户端接收答复的时间
*/
type ntpPackage struct {
	Originate   int64
	Receive     int64
	Transmit    int64
	Destination int64
}

/**
NTP报文的往返时延Delay=（T4-T1）-（T3-T2）=2秒。


t4-t1-t3+t2
Device A相对Device B的时间差offset=（（T2-T1）+（T3-T4））/2=1小时。
*/

func (n *ntpPackage) SysTimeNanoDiff() int64 {
	return (n.Receive - n.Originate - time.Now().UnixNano() + n.Transmit) / 2
}

func (n *ntpPackage) SysRoundTime() int64 {
	return time.Now().UnixNano() - n.Originate - n.Transmit + n.Receive
}

type ntpCmdFactory struct {
}

func (u ntpCmdFactory) GetCmd(cmd uint16) Command {

	switch cmd {
	case 1:
		return NetTimeStamp
	}

	return DefaultCommand{}
}

var NetTimeStamp = &netTimeStamp{}

// hi cmdcodec=0
// hi cmdres=1
type netTimeStamp struct {
	count int64
	avg   int64
}

func (d *netTimeStamp) Execute(ctx netty.InboundContext, message netty.Message) {

	//fmt.Println(message)

	//fmt.Println(".....ntp response  exec")

	// handler hi
	msg := message.([]byte)
	_, _, codec := DecodeHead(msg)

	if codec == 1 {

		t1 := binary.BigEndian.Uint64(msg[20:28])
		t2 := binary.BigEndian.Uint64(msg[28:36])
		t3 := binary.BigEndian.Uint64(msg[36:44])

		//ctx.Write(ntp)

		p := ntpPackage{Originate: int64(t1), Receive: int64(t2), Transmit: int64(t3)}

		diff := p.SysTimeNanoDiff()
		d.avg = (d.avg*d.count + diff) / (d.count + 1)
		d.count++
	}
}
