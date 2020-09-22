package restful

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "gw/event"
	. "gw/eventbus"
	"log"
)

type eventReqBody struct {
	Result   bool   `json:"result"`
	CaseInfo string `json:"case"`
	CaseCode int32  `json:"caseCode"`
	OrderId  string `json:"orderId"`
}

// map[carId:VIN30000S20003 orderId:1600139956 poi:map[]]
type srdReqBody struct {
	CarId   string                 `json:"carId"`
	OrderId string                 `json:"orderId"`
	Poi     map[string]interface{} `json:"poi"`
}

type srdResBody struct {
	Result     bool        `json:"result"`
	OrderId    string      `json:"orderId"`
	CaseInfo   string      `json:"case"`
	CaseCode   int32       `json:"caseCode"`
	NetworkId  interface{} `json:"networkId"`
	CockpitId  string      `json:"cockpitId"`
	DriverInfo interface{} `json:"driverInfo"`
}

// /erd/apply
func ErdApply(c *gin.Context) {
	var body eventReqBody
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println("parse error")
	}
	c.JSON(200, body)

}

//  no need
// /erd/stop
func ErdStop(c *gin.Context) {
	var body eventReqBody
	err := c.BindJSON(&body)

	log.Println("erd stop endpoint", body.OrderId)
	if err != nil {
		fmt.Println("parse error")
	}

	topic := "erd:stop:" + body.OrderId

	NewEvent(topic, body).Listen(NextFunc(func(e *Event) {

		GlobalBus.SubscribeOnce(topic, func(ee interface{}) {

			log.Println("erd stop sub ", body.OrderId, "sub", ee)

			if e.IsDisposed() {
				log.Println("already time out")
			} else {
				c.JSON(200, ee)
				e.Disposed()
			}
		})

		GlobalBus.Publish("erd:stop", gin.H{"orderId": body.OrderId})
	}), func(err error) {

		c.JSON(200, gin.H{"result": false,
			"cause":     "平台端处理超时",
			"causeCode": 504,
			"orderId":   body.OrderId,
		})
		c.JSON(200, srdResBody{Result: false, OrderId: body.OrderId})
	}, func() {
	})
}

// /erd/apply
func ErdAck(c *gin.Context) {

	var body gin.H
	err := c.BindJSON(&body)
	if err != nil {
		fmt.Println("parse error")
	}
	c.JSON(200, body)
}
