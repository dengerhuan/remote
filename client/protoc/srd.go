package protoc

import (
	"client/drivemanager"
	"encoding/json"
	"github.com/go-netty/go-netty"
	"log"
	"time"
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

}

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

		drivemanager.OrderId = orderId
		drivemanager.RdState = action

		drivemanager.StartTime = time.Now().UnixNano() / 1e9

		//order, ok := InstanceManager.Get(orderId)
	}
}
