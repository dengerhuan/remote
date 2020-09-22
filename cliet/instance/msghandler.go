package instance

import (
	"client/drivemanager"
	"client/protoc"
	"encoding/binary"
	"github.com/go-netty/go-netty"
	"log"
)

type MsgHandler struct {
}

func (MsgHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {

	msg := message.([]byte)

	switch msg[7] {

	case 0:
	// fmt.Println("keepalive")
	//ctx.Write([]byte(strconv.FormatInt(time.Now().UnixNano(),10)))

	case 2:
		log.Println(message)

		if len(msg) < 40 {
			log.Println(len(msg), msg)
			return
		}

		_, _, c := protoc.DecodeHead(msg)
		if c == 1 {

			//log.Println(msg[20:])
			drivemanager.ReadCommand(msg[20:])

			times := binary.BigEndian.Uint64(msg[32:])

			drivemanager.HH = int64(times)

			//
			//mide := int64(times) - drivemanager.SysTimeDiff
			//
			//log.Println((time.Now().UnixNano() - mide) / 1000 / 1000)
		}

	case 3:

		//fmt.Println(message)
	default:
		ctx.HandleRead(msg)
	}
}
