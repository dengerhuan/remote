package router

import "github.com/go-netty/go-netty"

type NetFilter interface {
	netty.OutboundHandler
	netty.InboundHandler
}

func NewNetFilter() NetFilter {
	return &netFilter{}
}

type netFilter struct {
}

func (netFilter) HandleRead(ctx netty.InboundContext, message netty.Message) {

}

// OutboundHandler defines an outbound handler

func (netFilter) HandleWrite(ctx netty.OutboundContext, message netty.Message) {

}
