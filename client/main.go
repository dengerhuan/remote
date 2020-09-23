package main

import (
	"client/.idea/netty/transport/udp"
	"client/authz"
	"client/instance"
	"client/protoc"
	"client/sys/ntp"
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/go-netty/go-netty/utils"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var boot = netty.NewBootstrap().Transport(udp.New())

func main() {
	log.SetFlags(log.Ldate | log.Llongfile)
	boot.ClientInitializer(func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(frame.PacketCodec(1024*6),
				exHandler{},

				protoc.CmdHandler{},
				instance.MsgHandler{},
			)
	})

	AfterInit()

	boot.Action(netty.WaitSignal(os.Interrupt, os.Kill))
	//------
}

// 断链 重连
func AfterInit() {

	con, err := boot.Connect("47.105.169.160:9090", nil)

	if err != nil {
		fmt.Println(err)
	}

	authz.SayHi(con)

	MockCmd(con)
	MockStat(con)

	authz.KeepAlive(con)
	utils.Assert(err)

	ticket := time.NewTicker(time.Second / 5)

	go func() {

		begin := 0

		for {

			if begin == 100 {
				ticket.Stop()
				runtime.Goexit()
			}
			select {
			case <-con.Context().Done():
				fmt.Println("channel done")
				ticket.Stop()
				runtime.Goexit()
			case <-ticket.C:
				con.Write(ntp.RequestNetTimeStamp())
			}
			begin++
		}

	}()
}

type exHandler struct{}

func (exHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	//AfterInit()

	fmt.Println(ex)
	if strings.Contains(ex.Error(), "connection refused") {
		ctx.Close(nil)
		log.Println("connection error ,reconnect after 2s")
		time.Sleep(time.Second * 2)
		AfterInit()
	}
}
