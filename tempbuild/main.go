package main

import (
	// framework
	cyako "github.com/Cyako/Cyako.go"

	// middlewares
	_ "github.com/Cyako/Cyako.go/kvstore"
	// _ "github.com/Cyako/Cyako.go/jsonbase"
	// _ "github.com/Cyako/Cyako.go/statistics"

	// processor codules
	_ "github.com/Cyako/example/module"

	// systec library
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

	err := c.Run(":12345", "/")
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
