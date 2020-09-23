package protoc

import (
	"encoding/json"
	"github.com/go-netty/go-netty"
	. "gw/eventbus"
	"log"
)

type erdApplyReq struct {
}

func (d erdApplyReq) Execute(ctx netty.InboundContext, message netty.Message) {

	//  get info from tsp
	log.Println("erd apply receive from server")
	msg := message.([]byte)

	var payload H
	_, l, codec := DecodeHead(msg)

	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)

		GlobalBus.Publish("erd:apply", payload)

	}
}

type erdStop struct {
}

func (d erdStop) Execute(ctx netty.InboundContext, message netty.Message) {
	log.Println("erd stop receive from server")

	log.Println()

	msg := message.([]byte)
	var payload H
	_, l, codec := DecodeHead(msg)

	log.Println(codec)
	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)

		id, ok := payload["orderId"]

		log.Println(payload)
		//
		if ok {
			GlobalBus.Publish("erd:stop:"+id.(string), payload)
		}
	}
}

type erdAckReq struct {
}

func (d erdAckReq) Execute(ctx netty.InboundContext, message netty.Message) {

	// request to  tsp
}
