package main

import (
	"fmt"

	"github.com/h4tr3d/iris/v6"
	"github.com/h4tr3d/iris/v6/adaptors/httprouter"
	"github.com/h4tr3d/iris/v6/adaptors/websocket"
)

func handleConnection(c websocket.Connection) {

	// Read events from browser
	c.On("chat", func(msg string) {

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)

		// Write message back to browser
		c.Emit("chat", msg)
	})

}

func main() {
	app := iris.New()
	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())

	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		Endpoint:        "/echo",
	})

	ws.OnConnection(handleConnection)

	// Adapt the websocket server.
	// you can adapt more than one of course.
	app.Adapt(ws)

	app.Get("/", func(ctx *iris.Context) {
		ctx.ServeFile("websockets.html", false) // second parameter: enable gzip?
	})

	app.Listen(":8080")
}
