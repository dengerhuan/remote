package main

import (
	"github.com/go-netty/go-netty"
	"lg/instance"
	"lg/protoc"
	"lg/rpc"
	"log"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

var (
	mockCmd [28]byte
)

func Mock(channel netty.Channel) {

	ticket := time.NewTicker(time.Second / 50)
	c := &rpc.Context{channel}

	time.Sleep(time.Second * 5)
	command := mockCmd[0:6]

	go func() {

		for {
			select {
			case <-ticket.C:
				steering := RandInt64(1, 65535)

				command[0] = byte(steering >> 8)
				command[1] = byte(steering)

				gas := RandInt64(1, 255)
				brake := RandInt64(1, 255)
				gear := RandInt64(1, 6)

				command[2] = byte(gas)
				command[3] = byte(brake)
				command[4] = byte(math.Pow(2, float64(gear)))
				command[5] = 1

				c.RenderBin(c.MsgHead(2, 0), command)

				//log.Println(command)

			case <-channel.Context().Done():
				log.Println("channel done")
				ticket.Stop()
				runtime.Goexit()
			}
		}

	}()

}

func ErdApply(channel netty.Channel) {

	if strings.EqualFold(instance.OrderId, "") {
		return
	}

	c := &rpc.Context{channel}
	c.RenderJson(c.CmdHead(3, 4), protoc.H{"orderId": instance.OrderId})

}

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
