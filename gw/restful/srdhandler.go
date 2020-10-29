package restful

import (
	"github.com/gin-gonic/gin"
	. "gw/event"
	. "gw/eventbus"
	"log"
)

// url /srd/apply //  check status
func SrdApply(c *gin.Context) {
	log.Println("srd apply endpoint handler")

	var body srdReqBody
	err := c.BindJSON(&body)
	log.Println(body)

	if err != nil {
		log.Println("parse error", err)
	}

	topic := "srd:apply:" + body.OrderId

	NewEvent(topic, body).Listen(NextFunc(func(e *Event) {


		//
		GlobalBus.SubscribeOnce(topic, func(ee interface{}) {
			log.Println("srd apply endpoint handler over")
			log.Println(ee)
			if e.IsDisposed() {
				log.Println("already time out")
			} else {
				c.JSON(200, ee)
				e.Disposed()
			}
		})
		GlobalBus.Publish("srd:apply", gin.H{"orderId": body.OrderId, "carId": body.CarId, "poi": body.Poi})
	}), func(err error) {
		log.Println(err)

		c.JSON(200, gin.H{"result": false,
			"cause":      "平台端处理超时",
			"causeCode":  504,
			"orderId":    body.OrderId,
			"networkId":  gin.H{},
			"cockpitId":  "",
			"driverInfo": gin.H{},
		})
	}, func() {
	})
}

// url /srd/start
func SrdStart(c *gin.Context) {
	var body eventReqBody
	err := c.BindJSON(&body)

	log.Println("srd start endpoint", body.OrderId)
	if err != nil {
		log.Println("parse error")
	}

	topic := "srd:start:" + body.OrderId

	log.Println(body)

	NewEvent(topic, body).Listen(NextFunc(func(e *Event) {

		GlobalBus.SubscribeOnce(topic, func(ee interface{}) {

			log.Println("srd start sub ", body.OrderId, "sub")

			log.Println(ee)
			if e.IsDisposed() {
				log.Println("already time out")
			} else {
				c.JSON(200, ee)
				e.Disposed()
			}
		})

		GlobalBus.Publish("srd:start", gin.H{"orderId": body.OrderId, "result": body.Result})
	}), func(err error) {
		log.Println("srd start sub ", body.OrderId, "sub", err)

		c.JSON(200, gin.H{"result": false,
			"cause":     "平台端处理超时",
			"causeCode": 504,
			"orderId":   body.OrderId,
		})
	}, func() {
	})
}
