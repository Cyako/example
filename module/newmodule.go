package module

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/kvstore"
	// "fmt"
)

type nmdep struct {
	KVStore *kvstore.KVStore
}
type NewModule struct {
	Dependences nmdep
	cyako.Module
}

func (n *NewModule) Get() {

}

func init() {
	var m = NewModule{
		Dependences: nmdep{
			KVStore: cyako.Svc["KVStore"].(*kvstore.KVStore),
		},
	}
	cyako.LoadModule(m)
}
