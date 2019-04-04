package internal

import (
	"os"
	"path"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// SetUp will initialize iris app
func SetUp() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	return app
}

// RunServer start the application
func RunServer(app *iris.Application) {
	app.Run(
		iris.Addr(os.Getenv("PORT")),
		iris.WithoutServerError(iris.ErrServerClosed))
}

// LoadEnvVars find and load environment variables
func LoadEnvVars() {
	base, _ := filepath.Abs(".")
	envPath := path.Join(base, "configs/env/development.env")
	godotenv.Load(envPath)
}
