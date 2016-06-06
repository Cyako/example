package module

import (
	cyako "github.com/Cyako/Cyako.go"
	"github.com/Cyako/Cyako.go/kvstore"

	"fmt"
)

type kvsdep struct {
	KVStore *kvstore.KVStore
}

type KVStoreTest struct {
	Dependences kvsdep
}

func (k KVStoreTest) Test(ctx *cyako.Ctx) {
	kvstore := k.Dependences.KVStore
	kvstore.SetWithScoped("KVStoreTest", "test", 123)
	v := kvstore.GetWithScoped("KVStoreTest", "test")
	fmt.Println(v)
}

func init() {
	var m = KVStoreTest{
		Dependences: kvsdep{
			KVStore: cyako.Svc["KVStore"].(*kvstore.KVStore),
		},
	}
	cyako.LoadModule(m)
}
