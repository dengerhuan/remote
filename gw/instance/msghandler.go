package instance

import (
	"encoding/json"
	"github.com/go-netty/go-netty"
	"log"
)

type MsgHandler struct {
}

func (m MsgHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {

	msg := message.([]byte)

	switch msg[7] {

	case 0:

		// keepalive
	case 1:

		//var payload map[string]interface{}
		//
		//if msg[9] == 0 {
		//
		//	err := mq.Produce("rdstat", strconv.FormatInt(time.Now().UnixNano()/1e6, 10), msg[20:])
		//	log.Println(err)
		//} else {
		//	err := mq.Produce("rdlog", strconv.FormatInt(time.Now().UnixNano()/1e6, 10), msg[20:])
		//	log.Println(err)
		//}
		//json.Unmarshal(msg[20:], &payload)
		//
		//log.Println(payload)
		//s.Broadcast(msg)
	}

	ctx.HandleRead(message)

}

func (s MsgHandler) Log(msg []byte) {

	var payload map[string]interface{}

	//log.Println(msg)
	//_, l, _ := DDecodeHead(msg)

	err := json.Unmarshal(msg[20:], &payload)

	if err != nil {
		log.Println(msg)
		log.Println(err)
	}

	log.Println(payload)
	//if err != nil {
	//	panic("")
	//}
	//if msg[9] == 0 {
	//	log.Println(payload)
	//} else {
	//	log.Println(payload)
	//}

}
