package main

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/go-netty/go-netty/utils"
	"gw/authz"
	"gw/instance"
	"gw/netty/transport/udp"
	"gw/protoc"
	"gw/restful"
	"gw/sys/ntp"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var boot = netty.NewBootstrap().Transport(udp.New())

var bizLogic = &instance.ClientLogic{}

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
	bizLogic.Listen()

	AfterInit()

	go restful.Start()

	boot.Action(netty.WaitSignal(os.Interrupt, os.Kill))
	//restful.Tt()

}

// 断链 重连
func AfterInit() {

	con, err := boot.Connect("47.105.169.160:9090", nil)

	//con, err := boot.Connect("0.0.0.0:9090", nil)

	if err != nil {
		fmt.Println(err)
	}

	authz.SayHi(con)

	authz.KeepAlive(con)
	utils.Assert(err)

	bizLogic.SetWrite(con)

	ticket := time.NewTicker(time.Second / 10)

	go func() {
		select {
		case <-con.Context().Done():
			fmt.Println("channel done")
			ticket.Stop()
			runtime.Goexit()
		case <-ticket.C:
			con.Write(ntp.RequestNetTimeStamp())
		}
	}()
}

type exHandler struct{}

func (exHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	//AfterInit()

	if strings.Contains(ex.Error(), "connection refused") {

		ctx.Close(nil)

		//
		//ctx.Channel().Close()
		fmt.Println("connection error ,reconnect after 2s")
		time.Sleep(time.Second * 2)
		AfterInit()
	}
}

//--------
