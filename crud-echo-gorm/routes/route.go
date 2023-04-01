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
	// 	Authenticated with JWT
	eUserJwt := eUser.Group("")
	eUserJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eUserJwt.GET("", controllers.GetUsersController)
	eUserJwt.GET("/:id", controllers.GetUserController)
	eUserJwt.PUT("/:id", controllers.UpdateUserController)
	eUserJwt.DELETE("/:id", controllers.DeleteUserController)

	// books routes
	eBook := e.Group("/books")
	// Authenticated with JWT
	eBook.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eBook.GET("", controllers.GetBooksController)
	eBook.POST("", controllers.CreateBookController)
	eBook.GET("/:id", controllers.GetBookController)
	eBook.PUT("/:id", controllers.UpdateBookController)
	eBook.DELETE("/:id", controllers.DeleteBookController)

	// blogs routes
	eBlog := e.Group("/blogs")
	eBlog.GET("", controllers.GetBlogsController)
	eBlog.GET("/:id", controllers.GetBlogController)
	// Authenticated with JWT
	eBlogJwt := eBlog.Group("")
	eBlogJwt.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	eBlogJwt.POST("", controllers.CreateBlogController)
	eBlogJwt.PUT("/:id", controllers.UpdateBlogController)
	eBlogJwt.DELETE("/:id", controllers.DeleteBlogController)

	return e
}
