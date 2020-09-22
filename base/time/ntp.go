package time

import (
	"encoding/binary"
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
	"strings"
	"time"
)

func NtpHandler(mode bool) netty.Handler {
	return &ntp{mode: mode}
}

type ntp struct {
	mode bool
	buf  [64]byte
}

type NtpStamp int64

/**
  Originate Timestamp       T1        客户端发送请求的时间
  Receive Timestamp        T2        服务器接收请求的时间
  Transmit Timestamp       T3        服务器答复时间
  Destination Timestamp     T4        客户端接收答复的时间
*/
type ntpPackage struct {

}

func (n *ntp) HandleRead(ctx netty.InboundContext, message netty.Message) {

	time.Now().UnixNano()

	msg := message.([]byte)

	start := binary.BigEndian.Uint32(msg[20 : 20+8])

	fmt.Println(start)
	/**
	  	reader := utils.MustToReader(message)
	  n, err := reader.Read(p.buffer)
	  utils.AssertIf(nil != err && io.EOF != err, "%v", err)

	  ctx.HandleRead(p.buffer[:n])
	*/

	// read text bytes
	textBytes := utils.MustToBytes(message)

	// convert from []byte to string
	sb := strings.Builder{}
	sb.Write(textBytes)

	// post text
	ctx.HandleRead(sb.String())
}

func (n *ntp) HandleWrite(ctx netty.OutboundContext, message netty.Message) {

	switch s := message.(type) {
	case string:
		ctx.HandleWrite(strings.NewReader(s))
	default:
		ctx.HandleWrite(message)
	}
}
