package restful

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "gw/eventbus"
	"log"
)

func init() {

	GlobalBus.SubscribeAsync("erd:apply", func(data map[string]interface{}) {
		log.Println(data)
		//
		//ErdApplyReq(data)
		//ErdApplyAck(data)

	}, false)
}

func Start() {

	r := gin.Default()

	r.POST("/srd/apply", SrdApply)
	r.POST("/srd/start", SrdStart)
	r.POST("/erd/apply", ErdApply) // to tsp
	r.POST("/erd/stop", ErdStop)
	r.POST("/erd/ack", ErdAck) // to tsp

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("server start error", err)
	} else {
		fmt.Println("serve and listen 8080")
	}

}
