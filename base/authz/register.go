package authz

import (
	"base/session"
	"github.com/go-netty/go-netty"
	"log"
)

var buf [21]byte

type RegisterHandler struct {
	Authenticated bool
	Session       *session.Session
}

func (s *RegisterHandler) HandleActive(ctx netty.ActiveContext) {
	s.Session = session.SessionManagerInst.SessionAt(ctx.Channel().ID())
	ctx.HandleActive()
}

func (s *RegisterHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {

	if s.Authenticated {
		ctx.HandleRead(message)
	} else if _, ok := s.Session.Attr("auth"); ok {
		s.Authenticated = true
		ctx.HandleRead(message)

	} else {
		//ctx.HandleRead(message)
		requestAuth(ctx.Channel())
	}
}

func requestAuth(channel netty.Channel) {

	res := buf[:]

	res[7] = 2
	res[9] = 1
	res[11] = 1
	res[15] = 1
	res[20] = 0

	channel.Write(res)

	log.Println(" request user for register")
}

func Auth(chlId int64, seq string) (vid Vehicle, ob bool) {

	vid, ok := GetVehicleBySeq(seq)
	//fmt.Println(seq)
	if ok {
		session := session.SessionManagerInst.SessionAt(chlId)
		session.SetAttr("auth", true)
		session.SetAttr("id", vid.vid)
	}
	return vid, ok
}
