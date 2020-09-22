package protoc

import (
	"encoding/json"
	"github.com/go-netty/go-netty"
	"lg/instance"
	"lg/rpc"
	"log"
	"strings"
	"time"
)

var OperationCmdFactory = operationCmdFactory{}

type operationCmdFactory struct {
}

func (u operationCmdFactory) GetCmd(cmd uint16) Command {

	switch cmd {

	case 5:
		return erdStopRes{}

	case 6:
		return erdStop{}

	}
	return erdStop{}
}

type erdStopRes struct {
}

func (d erdStopRes) Execute(ctx netty.InboundContext, message netty.Message) {
	log.Println("erd apply response from server")
	msg := message.([]byte)

	var payload H
	_, l, codec := DecodeHead(msg)
	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)
		//
		//id, ok := payload["orderId"]

		log.Println(payload)

	}
}

type erdStop struct {
}

func (d erdStop) Execute(ctx netty.InboundContext, message netty.Message) {
	log.Println("erd stop receive from server")
	msg := message.([]byte)

	var payload H
	_, l, codec := DecodeHead(msg)
	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)

		id, ok := payload["orderId"]

		log.Println(payload)

		if ok {
			instance.OrderId = id.(string)

			log.Println(id)
			time.Sleep(time.Second * 5)
			ErdApply(ctx.Channel())

		}
	}
}

func ErdApply(channel netty.Channel) {

	if strings.EqualFold(instance.OrderId, "") {
		return
	}

	c := &rpc.Context{channel}
	c.RenderJson(c.CmdHead(3, 4), H{"orderId": instance.OrderId})

}
