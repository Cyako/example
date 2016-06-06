package main

import (
	// framework
	cyako "github.com/Cyako/Cyako.go"

	// services
	// _ "github.com/Cyako/Cyako.go/kvstore"
	// _ "github.com/Cyako/Cyako.go/realtime"
	// _ "github.com/Cyako/Cyako.go/specvalue"ss
	// _ "github.com/Cyako/Cyako.go/jsonbase"
	// _ "github.com/Cyako/Cyako.go/statistics"

	// processor modules
	_ "github.com/Cyako/example/module"

	// system library
	"fmt"
	// "golang.org/x/net/websocket"
	// "net/http"
)

func main() {

	c := cyako.Ins()
	c.PrintLoadInfo()
	// c.PrintAPIDoc()
	c.CheckModule()
	fmt.Println()

	// SERVER
	fmt.Println(" Running...")
	fmt.Println()

	err := c.Run(":3000", "/")
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
