package module

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/kvstore"

	"fmt"
)

type kvsdep struct {
	KVStore *kvstore.KVStore
}

type KVStoreExample struct {
	Dependences kvsdep
}

func (k KVStoreExample) Example(ctx *cyako.Ctx) {
	kvstore := k.Dependences.KVStore
	kvstore.SetWithScoped("KVStoreExample", "test", 123)
	v := kvstore.GetWithScoped("KVStoreExample", "test")
	fmt.Println(v)
}

func init() {
	var m = KVStoreExample{
		Dependences: kvsdep{
			KVStore: cyako.Svc["KVStore"].(*kvstore.KVStore),
		},
	}
	cyako.LoadModule(m)
}
