package protoc

import (
	. "base/instance"
	"encoding/json"
	"github.com/go-netty/go-netty"
	"log"
)

type erdApply struct {
}

// get msg from cockpit redirect to vehicle and console
func (d erdApply) Execute(ctx netty.InboundContext, message netty.Message) {

	msg := message.([]byte)

	var payload H

	_, l, codec := DecodeHead(msg)

	if codec == 0 {

		json.Unmarshal(msg[20:20+l], &payload)
		orderId := payload["orderId"].(string)
		order, ok := InstanceManager.Get(orderId)

		var response H
		if !ok {
			response = H{"result": false,
				"cause":     "无效orderId",
				"causeCode": 601,
				"orderId":   orderId,
			}
			c := &Context{Write: ctx.Channel()}
			c.RenderJson(c.CmdHead(3, 5), response)
			return
		}

		log.Println("erd apply")

		log.Println(orderId, msg)

		response = H{"result": true,
			"cause":     "",
			"causeCode": 200,
			"orderId":   orderId,
		}

		order.GetConsole().Send(msg)
		order.GetVehicle().Send(msg)


		//
		//vg := order.GetVehicle().GetVg()
		//for _, logic := range vg {
		//	c := &Context{Write: logic.Ctx().Channel()}
		//	c.RenderJson(c.CmdHead(3, 5), response)
		//}

		//order.GetConsole().Send(msg)
		//order.GetVehicle().Send(msg)

		order.Stop()
	}
}

type erdStop struct {
}

// get from gw redirect to cockpit
func (d erdStop) Execute(ctx netty.InboundContext, message netty.Message) {

	msg := message.([]byte)

	var payload H

	c := &Context{Write: ctx.Channel()}
	_, l, codec := DecodeHead(msg)

	if codec == 0 {

		json.Unmarshal(msg[20:20+l], &payload)
		orderId := payload["orderId"].(string)
		order, ok := InstanceManager.Get(orderId)

		log.Println(order)

		if !ok {
			c.RenderJson(c.CmdHead(3, 7),
				H{"result": false,
					"cause":     "无效orderId",
					"causeCode": 601,
					"orderId":   orderId,
				})
			return
		}

		if !order.GetState() {
			log.Println("orderid", orderId, "is not working")

			c.RenderJson(c.CmdHead(3, 7),
				H{"result": false,
					"cause":     "order state is not working",
					"causeCode": 602,
					"orderId":   orderId,
				})
			return
		}
		//if order. {
		//
		//}

		order.GetCockpit().Send(message)

		c.RenderJson(c.CmdHead(3, 7),
			H{"result": true,
				"cause":     "",
				"causeCode": 200,
				"orderId":   orderId,
			})
	}
}

type erdAck struct {
}

func (d erdAck) Execute(ctx netty.InboundContext, message netty.Message) {

	msg := message.([]byte)

	var payload H
	_, l, codec := DecodeHead(msg)

	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)

		//fmt.Println(payload)
		//id, ok := payload["orderId"]
		////
		//if ok {
		//	GlobalBus.Publish("srd:apply:"+id.(string), message)
		//}
	}
}
