package main

import (
	"github.com/h4tr3d/iris/v6"
	"github.com/h4tr3d/iris/v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// Adapt the "httprouter", faster,
	// but it has limits on named path parameters' validation,
	// you can adapt "gorillamux" if you need regexp path validation!
	app.Adapt(httprouter.New())

	app.HandleFunc("GET", "/", func(ctx *iris.Context) {
		ctx.Writef("hello world\n")
	})

	app.Listen(":8080")
}
