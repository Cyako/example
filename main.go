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
	"net/http"
	"os"
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

	port := os.Getenv("PORT")
	c.Handle("/api/")
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
	// http.ListenAndServe(":"+port, http.FileServer(http.Dir("")))
	// err := c.Run(":"+port, "/api/")
	// if err != nil {
	// 	panic("ListenAndServe: " + err.Error())
	// }
}
