package instance

import (
	. "gw/eventbus"
	"gw/rpc"
	"io"
	"log"
)

type ClientLogic struct {
	writer io.Writer

	c *rpc.Context
}

func (l *ClientLogic) SetWrite(writer io.Writer) {
	l.writer = writer

	l.c = &rpc.Context{writer}

}

func (l *ClientLogic) Listen() {
	//
	GlobalBus.SubscribeAsync("srd:apply", func(data map[string]interface{}) {
		log.Println("srd,apply")
		l.c.RenderJson(l.c.CmdHead(3, 0), data)
	}, false)

	GlobalBus.SubscribeAsync("srd:start", func(data map[string]interface{}) {
		l.c.RenderJson(l.c.CmdHead(3, 2), data)
	}, false)

	//GlobalBus.SubscribeAsync("erd:apply", func(data map[string]interface{}) {
	//	l.c.RenderJson(l.c.CmdHead(3, 4), data)
	//}, false)

	GlobalBus.SubscribeAsync("erd:stop", func(data map[string]interface{}) {
		l.c.RenderJson(l.c.CmdHead(3, 6), data)
	}, false)

	//GlobalBus.SubscribeAsync("erd:ack", func(data map[string]interface{}) {
	//	l.c.RenderJson(l.c.CmdHead(3, 8), data)
	//}, false)
}
