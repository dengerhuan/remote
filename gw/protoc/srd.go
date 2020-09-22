package protoc

import (
	"encoding/json"
	"fmt"
	"github.com/go-netty/go-netty"
	. "gw/eventbus"
	"log"
)

type srdCheck struct {
}

func (d srdCheck) Execute(ctx netty.InboundContext, message netty.Message) {

	fmt.Println("srd apply protoc handler")
	msg := message.([]byte)

	var payload H
	_, l, codec := DecodeHead(msg)

	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)

		fmt.Println(payload)
		id, ok := payload["orderId"]
		//
		if ok {
			GlobalBus.Publish("srd:apply:"+id.(string), payload)
		}
	}
}

type srdStart struct {
}

func (d srdStart) Execute(ctx netty.InboundContext, message netty.Message) {

	log.Println("srd start  ", "proc srd")
	msg := message.([]byte)

	var payload H
	_, l, codec := DecodeHead(msg)

	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)


		id, ok := payload["orderId"]
		//
		if ok {
			GlobalBus.Publish("srd:start:"+id.(string), payload)
		}
	}
}
