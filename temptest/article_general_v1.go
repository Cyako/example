package xmodule

import (
	"fmt"
	cyako "github.com/Cyako/Cyako.go"
)

type Article struct{}

func (_ Article) GetIndexArticles(c *cyako.Ctx) {
	c.Set(&cyako.ParamConfig{Key: "num", Required: false, Default: "100", Echo: true})
	c.Set([]*cyako.ParamConfig{
		{Key: "p1", Default: "1"},
		{Key: "p2", Default: "2"},
		{Key: "p3", Default: "3"},
	})
	if c.Error.Fatal == nil {
		fmt.Println("No Fatal Error.")
		// p1 := c.Get("num")
		// p2 := c.Get("a")
		// c.Params["a"] = `456"""7"8""9`
		// fmt.Println(p1, p2)
		// fmt.Println(c.Params)
	}
}
func (_ Article) GetMessages(ctx *cyako.Ctx) {
	fmt.Println("M")
}
func init() {
	var m Article
	cyako.LoadModule(m)
}
