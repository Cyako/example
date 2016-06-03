package module

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/realtime"
)

type RealtimeExample struct {
	realtime *realtime.Realtime
}

func (r *RealtimeExample) Get() {

}

func init() {
	// cyako.LoadModule(&RealtimeExample{
	// 	realtime: &cyako.Ins().Middleware["Realtime"].(realtime.Realtime),
	// })
}
