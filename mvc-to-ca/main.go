package main

import (
	"mvc-to-ca/config"
	"mvc-to-ca/route"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	//init echo instance with routes
	e := route.NewRoute(db)

	e.Logger.Fatal(e.Start(":8000"))
}
