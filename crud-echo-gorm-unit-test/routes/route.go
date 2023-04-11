package routes

import (
	"crud-echo-gorm/constants"
	"crud-echo-gorm/controllers"
	m "crud-echo-gorm/middlewares"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	//		Authentication with Basic Auth
	// eAuthBasic := e.Group("/auth")
	// eAuthBasic.Use(mid.BasicAuth(m.BasicAuthDB))
	// //users routes
	// eAuthBasic.GET("/users", controllers.GetUsersController)
	// eAuthBasic.GET("/users/:id", controllers.GetUserController)
	// eAuthBasic.POST("/users", controllers.CreateUserController)
	// eAuthBasic.PUT("/users/:id", controllers.UpdateUserController)
	// eAuthBasic.DELETE("/users/:id", controllers.DeleteUserController)

	// users routes
	eUser := e.Group("/users")
	eUser.POST("", controllers.CreateUserController)
	eUser.POST("/login", controllers.LoginUserController)
	// 	users routes authenticated with JWT
	eUserJwt := eUser.Group("")
	eUserJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUserJwt.GET("", controllers.GetUsersController)
	eUserJwt.GET("/:id", controllers.GetUserController)
	eUserJwt.PUT("/:id", controllers.UpdateUserController)
	eUserJwt.DELETE("/:id", controllers.DeleteUserController)

	// books routes
	eBookJWT := e.Group("/books")
	// books routes Authenticated with JWT
	eBookJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eBookJWT.GET("", controllers.GetBooksController)
	eBookJWT.POST("", controllers.CreateBookController)
	eBookJWT.GET("/:id", controllers.GetBookController)
	eBookJWT.PUT("/:id", controllers.UpdateBookController)
	eBookJWT.DELETE("/:id", controllers.DeleteBookController)

	// blogs routes
	// e.GET("/blogs",controllers.GetBlogsController)

	return e
}
