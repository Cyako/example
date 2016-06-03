package main

import (
	cyako "github.com/Cyako/Cyako.go"
	_ "github.com/Cyako/Cyako.go/realtime"

	_ "github.com/Cyako/example/module"
)

func main() {
	c := cyako.Ins()
	c.Run(":3000", "/")
}
