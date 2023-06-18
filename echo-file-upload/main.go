package main

import (
	"file-upload/controllers"
	filemanager "file-upload/lib/file_manager"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := filemanager.InitAssetsFile()
	if err != nil {
		log.Fatal(err)
	}
	
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", controllers.Upload)
	e.GET("/index", controllers.Index)

	e.Logger.Fatal(e.Start(":8000"))
}
