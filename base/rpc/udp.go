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
	fmt.Println(ex)

}

//
//type loggerHandler struct {
//	buf [1024 * 4]byte
//}
//
//func (l *loggerHandler) HandleActive(ctx netty.ActiveContext) {
//	fmt.Println("go-netty:", "->", "active:", ctx.Channel().RemoteAddr())
//	// write welcome message
//
//	var _buf = l.buf[:]
//
//	//fmt.Printf("%T",l.buf)
//
//	ctx.HandleActive()
//
//	//id := ctx.Channel().ID()
//	//fmt.Println(session.ManagerInst.Size())
//	//fmt.Println(session.ManagerInst.SessionAt(id))
//	var i int32 = 0
//	go func() {
//		for i < 10 {
//
//			//fmt.Println(i)
//			_buf[0] = byte(i)
//			i++
//
//			//fmt.Println(i)
//			ctx.Write(_buf)
//		}
//	}()
//
//	//ctx.Write([]byte("Hello I'm " + "go-netty"))
//}
//
//func (l *loggerHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
//
//	fmt.Println("logger handler")
//	ctx.HandleRead(string(message.([]byte)))
//}
//
//func (loggerHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
//	fmt.Println("go-netty:", "->", "inactive:", ctx.Channel().RemoteAddr(), ex)
//	// disconnected，the default processing is to close the connection
//	ctx.HandleInactive(ex)
//}
//
//type upperHandler struct{}
//
//func (upperHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
//	// text to upper case
//
//	if message.(string) == "9" {
//
//		fmt.Println("ssss")
//		//ctx.Close(nil)
//
//	}
//
//	fmt.Println(message)
//
//}
