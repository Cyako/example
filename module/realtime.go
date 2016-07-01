package module

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/realtime"
)

type rtdev struct {
	Realtime *realtime.Realtime
}

type RealtimeExample struct {
	Dependences rtdev
}

func (r RealtimeExample) JoinChatRoom(ctx *cyako.Ctx) {
	realtime := r.Dependences.Realtime
	// realtime.AddListener("chatroom", ctx.Conn, ctx.Id, ctx.Method)
	realtime.AddListenerDefault("chatroom", ctx)
}

func (r RealtimeExample) SendChatMessage(ctx *cyako.Ctx) {
	realtime := r.Dependences.Realtime
	ctx.Set(&cyako.ParamConfig{Key: "message", Required: true})
	res := &cyako.Res{}
	res.Init()
	res.Params["message"] = ctx.Params["message"]
	realtime.Send("chatroom", res)
}

func init() {
	var m = RealtimeExample{
		Dependences: rtdev{
			Realtime: cyako.Svc["Realtime"].(*realtime.Realtime),
		},
	}
	cyako.LoadModule(m)
}
