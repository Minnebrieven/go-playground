package controller

import (
	"net/http"
	"static-api-crud-echo/config"
	"static-api-crud-echo/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

var users []model.User

// var users = []User{
// 	{
// 		ID:       1,
// 		Name:     "Fikri",
// 		Email:    "fikri@gmail.com",
// 		Password: "ajksidohtmpqsmsalk",
// 	},
// 	{
// 		ID:       2,
// 		Name:     "Andy",
// 		Email:    "Andy@gmail.com",
// 		Password: "sadoiqwesjirpa",
// 	},
// 	{
// 		ID:       3,
// 		Name:     "Ceb",
// 		Email:    "ceb@gmail.com",
// 		Password: "sdavbtashjyu",
// 	},
// }

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {

	resp := map[string]interface{}{
		"messages": "users not found",
		"users":    nil,
	}

	id, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		resp["messages"] = errorId
		return c.JSON(http.StatusBadRequest, resp)
	}

	var user = model.User{}
	result := config.DB.Where("id = ?", id).First(&user)

	if result.Error != nil {
		resp["messages"] = result.Error
		return c.JSON(http.StatusOK, resp)
	}

	resp["messages"] = "success find user by id"
	resp["users"] = user

	return c.JSON(http.StatusOK, resp)
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	resp := map[string]interface{}{
		"messages": "users not found",
		"users":    nil,
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp["messages"] = err
		return c.JSON(http.StatusBadRequest, resp)
	}

	for idx := range users {
		if users[idx].ID == id {
			resp["users"] = users[idx]

			copy(users[idx:], users[idx+1:])   // Shift a[i+1:] left one index.
			users[len(users)-1] = model.User{} // Erase last element (write zero value).
			users = users[:len(users)-1]

			resp["messages"] = "user deleted"

		}
	}

	return c.JSON(http.StatusOK, resp)
}

// update user by id
func UpdateUserController(c echo.Context) error {
	resp := map[string]interface{}{
		"messages": "users not found",
		"users":    nil,
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp["messages"] = err
		return c.JSON(http.StatusBadRequest, resp)
	}

	for idx := range users {
		if users[idx].ID == id {
			users[idx].Name = c.FormValue("name")
			users[idx].Email = c.FormValue("email")
			users[idx].Password = c.FormValue("password")

			resp["messages"] = "success update users"
			resp["users"] = users[idx]

			break
		}
	}

	return c.JSON(http.StatusOK, resp)
}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}
