package instance

import (
	"fmt"
	"github.com/go-netty/go-netty"
	"log"
)

var OrderId string = ""

type MsgHandler struct {
}

func (MsgHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {

	msg := message.([]byte)

	switch msg[7] {

	case 0:
		// fmt.Println("keepalive")
		//ctx.Write([]byte(strconv.FormatInt(time.Now().UnixNano(),10)))
	case 3:

		fmt.Println(message)

		log.Println(msg[20])
		log.Println(msg[21])
		log.Println("set category")

	default:

		fmt.Println(message)
		ctx.HandleRead(msg)
	}
}
