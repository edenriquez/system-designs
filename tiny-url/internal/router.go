package internal

import (
	"github.com/kataras/iris"
)

// RouterSetUp shoulddeclare routes
func RouterSetUp(app *iris.Application) {
	app.Handle("GET", "/health-check", healthCheckHandler)
	app.Handle("POST", "/new/url", urlEncodeHandler)
}
