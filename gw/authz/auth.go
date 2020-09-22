package authz

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"gw/rpc"
	"gw/tools"
	"runtime"
	"strconv"
	"time"
)

var (
	keepalive [28]byte
)

func SayHi(channel netty.Channel) {

	context := &rpc.Context{Write:channel}

	chead := context.CmdHead(2, 0)

	context.RenderString(chead, tools.GenDeviceIdByHardDisk())

}

func Register(channel netty.Channel) {

	context := &rpc.Context{channel}

	chead := context.CmdHead(2, 2)

	context.RenderString(chead, tools.GenDeviceIdByHardDisk())
	//head := re[:20]
	//head[7] = 2
	//head[9] = 2
	//head[11] = 2
	//head[15] = 32
	//packet := append(head, []byte(tools.GenDeviceIdByHardDisk())...)
	//
	fmt.Println("register server with hex hard card id ")
	//
	//channel.Write(packet)
}

func KeepAlive(channel netty.Channel) {
	fmt.Println("keepalive request")
	msg := keepalive[:20]
	msg[6] = 1
	msg[7] = 0
	ticket := time.NewTicker(time.Second * 1)

	go func() {

		for {
			select {
			case <-ticket.C:
				s := strconv.FormatInt(time.Now().UnixNano(), 10)
				channel.Write(append(msg, []byte(s)...))

			case <-channel.Context().Done():
				fmt.Println("channel done")
				ticket.Stop()
				runtime.Goexit()
			}
		}

	}()

}
