package main

import (
	"crud-echo-gorm/config"
	"crud-echo-gorm/routes"
)

func init() {
	// initiate db connection
	config.InitDB()
}

func main() {
	// create a new echo instance
	e := routes.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
