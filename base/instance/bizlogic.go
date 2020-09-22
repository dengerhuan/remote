package instance

import (
	"encoding/binary"
	"fmt"
	"github.com/go-netty/go-netty"
	"log"
	"sync"
	"time"
)

//----- download ok

type RdLogic struct {
	mutex         sync.RWMutex
	IdleTime      time.Duration
	lastWriteTime int64
	handlerCtx    netty.HandlerContext
	id            string

	adminState bool
	rdState    bool

	state uint8 //  status 0 unlock  1 lock  2 working

	initAttach bool
	logicType  uint8 // 21 类型  --console 2 --车 0 - 台驾--1
	category   uint8 //20  能力
	vg         []*RdLogic
}

func (s *RdLogic) Ctx() netty.HandlerContext {
	return s.handlerCtx
}

func (s *RdLogic) GetVg() []*RdLogic {
	return s.vg
}

func (s *RdLogic) GetId() string {
	return s.id
}

func (s *RdLogic) Lock() {
	s.state = 1
}

func (s *RdLogic) Unlock() {
	s.state = 0
}

func (s *RdLogic) Acquire() bool {

	if s.state == 0 {
		s.state = 1
		return true
	}

	return false

}

func (s *RdLogic) SetState(state uint8) {
	s.state = state
}
func (s *RdLogic) Send(msg netty.Message) {
	s.mutex.Lock()
	s.handlerCtx.Write(msg)
	s.mutex.Unlock()
}

func (s *RdLogic) ChannelValid() bool {
	return time.Now().UnixNano()-s.lastWriteTime < 1e9*2
}
func (s *RdLogic) setCategory(message netty.Message) error {

	if s.initAttach {
		return nil
	}
	s.initAttach = true
	msg := message.([]byte)
	l := binary.BigEndian.Uint32(msg[12:16])

	if msg[11] != 1 {
		return fmt.Errorf("unsupport category codec")
	}

	tmpId := string(msg[22 : 22+l])

	s.id = tmpId

	s.category = msg[20]
	s.logicType = msg[21]

	//log.Println(s)

	group := registerByLogic(s.logicType)

	group.Register(s)

	return nil
}

func (s *RdLogic) HandleActive(ctx netty.ActiveContext) {
	log.Println("log active")
	s.handlerCtx = ctx
	ctx.HandleActive()
}

func (s *RdLogic) HandleRead(ctx netty.InboundContext, message netty.Message) {

	s.lastWriteTime = time.Now().UnixNano()
	msg := message.([]byte)

	switch msg[7] {

	case 0:
		// keep alive
	case 1:

		// log
		if s.logicType == 0 {
			packet := make([]byte, len(msg))
			copy(packet, msg)
			s.Broadcast(packet)
		}
	case 3:
		err := s.setCategory(message)
		if nil != err {
			log.Println(err)
		}

	case 2:
		l := len(msg)
		packet := make([]byte, l+14)
		copy(packet, msg)
		binary.BigEndian.PutUint64(packet[l+6:], uint64(time.Now().UnixNano()))
		s.Broadcast(packet)

	default:
		ctx.HandleRead(msg)
	}
}

func (s *RdLogic) Broadcast(message netty.Message) {

	s.mutex.RLock()
	defer s.mutex.RUnlock()
	for _, rd := range s.vg {
		rd.handlerCtx.Write(message)
	}
}

func (s *RdLogic) BroadcastIf(message netty.Message, fn func(logic *RdLogic) bool) {

	for _, rd := range s.vg {
		if fn(rd) {
			rd.handlerCtx.Write(message)
		}
	}
}

//if msg[9] == 1 {
//	//fmt.Println(msg)
//	s.Broadcast(msg)
//} else {
//
//	var payload map[string]interface{}
//
//	_, l, _ := DDecodeHead(msg)
//
//	err := json.Unmarshal(msg[20:20+l], &payload)
//
//	//fmt.Println(err)
//	//
//	//fmt.Println(payload)
//	//	stat
//	//	stat := H{"status": []map{}{carStat}, "time": time.Now().UnixNano() / 1e6}
//
//	status := payload["status"].([]interface{})
//	for _, _statu := range status {
//		statu := _statu.(map[string]interface{})
//		statu["carId"] = s.id
//
//	}
//
//	jsonBytes, err := json.Marshal(map[string]interface{}{
//		"status": status,
//	})
//	if err != nil {
//		return
//	}
//
//	h := make([]byte, 20)
//	copy(h, msg[0:20])
//	binary.BigEndian.PutUint32(h[12:16], uint32(len(jsonBytes)))
//
//	s.Broadcast(append(h, jsonBytes...))
//
//}

//log.Println("-------------")
//}

func DDecodeHead(msg []byte) (uint8, uint32, uint8) {
	flag := msg[10]
	l := binary.BigEndian.Uint32(msg[12:16])
	codec := msg[11]
	return flag, l, codec
}
