package rpc

import (
	"base/authz"
	"base/instance"
	"base/netty/transport/udp"
	"base/protoc"
	"base/session"
	"fmt"
	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/frame"
	"log"
	"os"
	"time"
)

// three kind method to insert handler
// server shared handler / channel shared handler / message shared handler

// cmd --register---logic--session

// --  if no register over it tx in cmd
func UDP() {

	// child pipeline initializer
	var childPipelineInitializer = func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(frame.PacketCodec(1024*4)).
			AddLast(session.SessionManagerInst,
				protoc.CmdHandler{},
				&authz.RegisterHandler{Authenticated: false},
				&instance.RdLogic{IdleTime: time.Second * 2}, exHandler{}, )
	}

	// new go-netty bootstrap
	biit := netty.NewBootstrap().
		// configure the child pipeline initializer
		ChildInitializer(childPipelineInitializer).
		// configure the transport protocol
		Transport(udp.New()).
		// configure the listening address
		Listen("0.0.0.0:9090").
		// waiting for exit signal
		Action(netty.WaitSignal(os.Interrupt, os.Kill)).
		// print exit message
		Action(func(bs netty.Bootstrap) {
			fmt.Println("server exited")
		})

	select {
	case <-biit.Context().Done():
		fmt.Println(biit.Context().Err())
	}

}

type exHandler struct{}

func (exHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	log.Println(ex)

}
