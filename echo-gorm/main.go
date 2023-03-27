package main

import (
	"static-api-crud-echo/config"
	"static-api-crud-echo/route"
)

// ---------------------------------------------------
func main() {
	config.InitDB()
	e := route.New()
	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
