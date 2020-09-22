package authz

import (
	"client/tools"
	"fmt"
	"github.com/go-netty/go-netty"
	"runtime"
	"strconv"
	"time"
)

var (
	keepalive [28]byte
)

func SayHi(channel netty.Channel) {
	hi := make([]byte,100)
	head := hi[:20]

	head[7] = 2

	head[11] = 2

	head[15] = 32

	packet := append(head, []byte(tools.GenDeviceIdByHardDisk())...)
	channel.Write(packet)

	fmt.Println("check status if device register")

}

func Register(channel netty.Channel) {

	hi := make([]byte,100)

	head := hi[:20]
	head[7] = 2
	head[9] = 2
	head[11] = 2
	head[15] = 32
	packet := append(head, []byte(tools.GenDeviceIdByHardDisk())...)

	fmt.Println(tools.GenDeviceIdByHardDisk())

	fmt.Println("register server with hex hard card id ")

	channel.Write(packet)
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
