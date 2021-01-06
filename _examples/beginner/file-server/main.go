package main

import (
	"github.com/h4tr3d/iris"
	"github.com/h4tr3d/iris/adaptors/httprouter"
)

func main() {
	app := iris.New()
	app.Adapt(httprouter.New())
	// first parameter is the request path
	// second is the operating system directory
	app.StaticWeb("/static", "./assets")

	app.Listen(":8080")
}
