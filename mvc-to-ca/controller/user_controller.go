package controller

import (
	"mvc-to-ca/model"
	"mvc-to-ca/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController interface{}

type userController struct {
	userUseCase usecase.UserUsecase
}

func NewUserController(userUC usecase.UserUsecase) *userController {
	return &userController{userUC}
}

func (u *userController) GetAllUsers(c echo.Context) error {
	users, err := u.userUseCase.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   users,
	})
}

func (u *userController) GetUserByID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "id parameter must be valid",
		})
	}

	var user = model.User{ID: uint(userID)}
	user, err = u.userUseCase.GetUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   user,
	})
}

func (u *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	token, err := u.userUseCase.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"token":  token,
	})
}

func (u *userController) CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}
	if err := u.userUseCase.CreateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "success",
		"data":   user,
	})
}
