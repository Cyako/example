package module

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/realtime"

	"fmt"
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
	fmt.Println("Send:", res.Params)
	fmt.Println(realtime)
	realtime.Send("chatroom", res)
	fmt.Println("Send finish.")
}

func init() {
	var m = RealtimeExample{
		Dependences: rtdev{
			Realtime: cyako.Svc["Realtime"].(*realtime.Realtime),
		},
	}
	cyako.LoadModule(m)
}
