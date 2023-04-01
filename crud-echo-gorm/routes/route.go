package routes

import (
	"crud-echo-gorm/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	//users routes
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	//books routes
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)

	//blogs routes
	e.GET("/blogs", controllers.GetBlogsController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)

	return e
}
