package main

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/frame"
	"github.com/go-netty/go-netty/utils"
	"lg/authz"
	_ "lg/media"
	"lg/netty/transport/udp"
	"lg/protoc"
	"log"
	"os"
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
			)
	})

	//con, err := boot.Connect("0.0.0.0:9090", nil)
	//
	//authz.SayHi(con)
	//authz.SayHi(con)
	//authz.KeepAlive(con)
	//
	//utils.Assert(err)

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

	authz.KeepAlive(con)

	//go func() {
	//	time.Sleep(time.Second * 20)
	//	ErdApply(con)
	//}()
	utils.Assert(err)
	Mock(con)

	//ticket := time.NewTicker(time.Second / 10)
	//
	//go func() {
	//
	//
	//	select {
	//	case <-con.Context().Done():
	//		log.Println("channel done")
	//		ticket.Stop()
	//		runtime.Goexit()
	//	case <-ticket.C:
	//
	//		con.Write(ntp.RequestNetTimeStamp())
	//	}
	//
	//}()
}

type exHandler struct{}

func (exHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	//AfterInit()

	log.Println(ex)

	if strings.Contains(ex.Error(), "connection refused") {
		ctx.Close(nil)
		fmt.Println("connection error ,reconnect after 2s")
		time.Sleep(time.Second * 2)
		AfterInit()
	}
}
