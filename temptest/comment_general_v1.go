package xmodule

import (
	"fmt"
	cyako "github.com/Cyako/Cyako.go"
)

type Comment struct{}

func (m Comment) GetIndexComments(ctx *cyako.Ctx) {
	fmt.Println("C")
}
func (m Comment) GetMessages(ctx *cyako.Ctx) {
	fmt.Println("M")
}
func init() {
	var m Comment
	cyako.LoadModule(m)
}
