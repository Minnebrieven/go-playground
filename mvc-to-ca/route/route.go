package route

import (
	"mvc-to-ca/constants"
	"mvc-to-ca/controller"
	m "mvc-to-ca/middleware"
	"mvc-to-ca/repository"
	"mvc-to-ca/usecase"

	mid "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB) *echo.Echo {
	e := echo.New()

	//REPOSITORIES
	userRepository := repository.NewUserRepository(db)

	//SERVICES or USECASES
	userService := usecase.NewUserUseCase(userRepository)

	//CONTROLLERS
	userController := controller.NewUserController(userService)

	//ROUTES

	//users Routes
	m.LogMiddleware(e)

	e.POST("/login", userController.Login)
	e.POST("/users", userController.CreateUser)

	usersJWT := e.Group("/users")
	usersJWT.Use(mid.JWT([]byte(constants.SECRET_JWT)))
	usersJWT.GET("", userController.GetAllUsers)
	usersJWT.GET("/:id", userController.GetUserByID)

	return e
}
