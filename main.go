package main

import (
	"github.com/yamato0204/go-test.git/app/registry"
	"github.com/yamato0204/go-test.git/app/router"
)


func main() {
	Handler := registry.NewRsolver()
	

	e := router.Route(Handler)

	e.Logger.Fatal(e.Start(":8080"))

}