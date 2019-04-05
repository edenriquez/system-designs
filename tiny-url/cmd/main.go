package main

import (
	internal "github.com/edenriquez/system-designs/tiny-url/internal"
	database "github.com/edenriquez/system-designs/tiny-url/internal/models"
)

func main() {
	// load environment variables
	internal.LoadEnvVars()
	// set up database connection and table migration
	database.SetUpDB()
	// instance iris app
	app := internal.SetUp()
	// register routes
	internal.RouterSetUp(app)
	// start server
	internal.RunServer(app)
}
