package protoc

import (
	. "base/instance"
	"encoding/json"
	"github.com/go-netty/go-netty"
	"log"
)

/**
payload
map[carId:12345678901234567 orderId:2731489313 poi:map[destination:XXX路303号楼 latitude:125.3249352 longitude:43.8593245]]

test data
{
"carid":"12345678901234567",
"orderId":"2731489313",
      "poi":{
            "latitude": 125.3249352,
            "longitude": 43.8593245,
            "destination":"XXX路303号楼"
  }
 }

response "{
“result”:true,
“cause”:“”，
“causeCode”:200,
“orderId”:“2731489313”,
“networkid”：{
    “CC”:”460“，
    “NC”:”000”,
    “MC”:“000000001”
 }，
“cockpitid”:“0001000001”，
“driverInfo”:{
    “name”:“dennis”,
    “id”:“0001000021”，
    “tel”: “8613911433443 ”
}
   }"


"


Http Body:
  {
“result”:false,
“cause”: “车辆状态NOK”，
“causeCode”: 611,
“orderId”:“2731489313”,
“networkid”：{}，
“cockpitid”:””,
“driverInfo”:{}
   }
"
*/

type srdCheck struct {
}

func (d srdCheck) Execute(ctx netty.InboundContext, message netty.Message) {

	log.Println(message)

	msg := message.([]byte)
	c := &Context{Write: ctx.Channel()}
	var payload H

	_, l, codec := DecodeHead(msg)
	if codec == 0 {
		err := json.Unmarshal(msg[20:20+l], &payload)

		if err != nil {
			log.Println(err)
		}

		log.Println(payload)

		orderId := payload["orderId"].(string)
		carId := payload["carId"].(string)

		order, ok := InstanceManager.Get(orderId)

		log.Println(carId)
		log.Println(orderId)

		log.Println(ok)
		if ok {

			if order.IsCancelled() {
				c.RenderJson(c.CmdHead(3, 1), H{"result": true,
					"cause":      "cancelled order",
					"causeCode":  613,
					"orderId":    orderId,
					"networkId":  H{},
					"cockpitId":  "",
					"driverInfo": H{},
				})

				return
			}

			c.RenderJson(c.CmdHead(3, 1), H{"result": true,
				"cause":      "",
				"causeCode":  200,
				"orderId":    orderId,
				"networkId":  H{"CC": "460", "NC": "00", "MC": "000000001"},
				"cockpitId":  order.GetCockpit().GetId(),
				"driverInfo": H{"name": "dennis", "id": "0001000021", "tel": "8613911433443"},
			})
			return
		}

		order = InstanceManager.Create(orderId)

		log.Println("create order", order)
		err = order.SetVehicle(carId)

		if err != nil {

			c.RenderJson(c.CmdHead(3, 1), H{"result": false,
				"cause":      "车辆状态NOK",
				"causeCode":  611,
				"orderId":    orderId,
				"networkId":  H{"CC": "460", "NC": "00", "MC": "000000001"},
				"cockpitId":  "",
				"driverInfo": H{},
			})
			return
		}

		log.Println("set vehicle ok")
		err = order.SetCockpit("")

		log.Println("set cockpit ok")

		if err != nil {
			c.RenderJson(c.CmdHead(3, 1), H{"result": false,
				"cause":      "台驾状态NOK",
				"causeCode":  612,
				"orderId":    orderId,
				"networkId":  H{"CC": "460", "NC": "00", "MC": "000000001"},
				"cockpitId":  "",
				"driverInfo": H{},
			})
			return
		}

		log.Println("set cockpit complete")

		c.RenderJson(c.CmdHead(3, 1), H{"result": true,
			"cause":      "",
			"causeCode":  200,
			"orderId":    orderId,
			"networkId":  H{"CC": "460", "NC": "00", "MC": "000000001"},
			"cockpitId":  order.GetCockpitId(),
			"driverInfo": H{"name": "dennis", "id": "0001000021", "tel": "8613911433443"},
		})

	}
}

type srdStart struct {
}

func (d srdStart) Execute(ctx netty.InboundContext, message netty.Message) {

	log.Println("srd start")

	msg := message.([]byte)

	c := &Context{Write: ctx.Channel()}

	var payload H
	_, l, codec := DecodeHead(msg)

	if codec == 0 {
		json.Unmarshal(msg[20:20+l], &payload)

		log.Println(payload)

		orderId := payload["orderId"].(string)
		action := payload["result"].(bool)
		order, ok := InstanceManager.Get(orderId)

		if !ok {
			log.Println(action)
			c.RenderJson(c.CmdHead(3, 3),
				H{"result": false,
					"cause":     "无效orderId",
					"causeCode": 601,
					"orderId":   orderId,
				})
			return
		}
		if action {

			order.Start()
			order.GetVehicle().Send(msg)
			c.RenderJson(c.CmdHead(3, 3),
				H{"result": true,
					"cause":     "",
					"causeCode": 200,
					"orderId":   orderId,
				})

		} else {

			order.Stop()
			order.GetVehicle().Send(msg)
			c.RenderJson(c.CmdHead(3, 3),
				H{"result": false,
					"cause":     "用户取消",
					"causeCode": 621,
					"orderId":   orderId,
				})

		}
	}

}
