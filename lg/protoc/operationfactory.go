package protoc

import (
	"encoding/json"
	"github.com/go-netty/go-netty"
	. "lg/eventbus"
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

	case 2:
		return srdStart{}
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

			//
			// stop  rtmp

			GlobalBus.Publish("rdstop", id)
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

//----

type srdStart struct {
}

func (d srdStart) Execute(ctx netty.InboundContext, message netty.Message) {
	log.Println("srd start")

	msg := message.([]byte)

	//c := &Context{Write: ctx.Channel()}

	var payload map[string]interface{}
	_, l, codec := DecodeHead(msg)

	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)

		log.Println(payload)

		orderId := payload["orderId"].(string)
		action := payload["result"].(bool)

		log.Println(orderId)
		log.Println(action)

		log.Println("start")

		GlobalBus.Publish("rdstart", orderId)
		//
		//order, ok := InstanceManager.Get(orderId)
	}
}
