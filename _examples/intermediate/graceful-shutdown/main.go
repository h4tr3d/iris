package main

import (
	"context"
	"time"

	"github.com/h4tr3d/iris/v6"
	"github.com/h4tr3d/iris/v6/adaptors/httprouter"
)

func main() {
	app := iris.New()
	// output startup banner and error logs on os.Stdout
	app.Adapt(iris.DevLogger())
	// set the router, you can choose gorillamux too
	app.Adapt(httprouter.New())

	app.Get("/hi", func(ctx *iris.Context) {
		ctx.HTML(iris.StatusOK, " <h1>hi, I just exist in order to see if the server is closed</h1>")
	})

	app.Adapt(iris.EventPolicy{
		// Interrupt Event means when control+C pressed on terminal.
		Interrupted: func(*iris.Framework) {
			// shut down gracefully, but wait 5 seconds the maximum before closed
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			app.Shutdown(ctx)
		},
	})

	app.Listen(":8080")
}
