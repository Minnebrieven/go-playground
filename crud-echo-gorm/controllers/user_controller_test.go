package controllers

import (
	"crud-echo-gorm/config"
	"crud-echo-gorm/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.InitDBTest()
	if err := CreateUserMocking(); err != nil {
		panic(err)
	}
}

func CreateUserMocking() error {
	user := models.User{
		Name:     "alterra",
		Email:    "alterra@academy.id",
		Password: "123",
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func InitEchoTest() *echo.Echo {
	e := echo.New()

	return e
}

func TestGetUsersController(t *testing.T) {
	testCases := []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Positive Case Get All Users",
			path:                 "/users",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"messages\":\"success get all users\",\"users\":[",
		},
		// {
		// 	name:                 "Negative Case wrong path",
		// 	path:                 "/user",
		// 	expectStatus:         http.StatusNotFound,
		// 	expectBodyStartsWith: "{\"message\":\"Not Found",
		// },
	}
	e := InitEchoTest()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		//Assertion
		if assert.NoError(t, GetUsersController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestGetUserController(t *testing.T) {
	testCases := []struct {
		name                 string
		path                 string
		userID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Positive Case Get User By ID",
			path:                 "/users/:id",
			userID:               "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"messages\":\"success get user\",\"users\":",
		},
		{
			name:                 "Negative Case Get User By ID not Exist",
			path:                 "/users/:id",
			userID:               "9999",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"messages\":\"record not found",
		},
	}

	e := InitEchoTest()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

		//Assertion
		if assert.NoError(t, GetUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestCreateUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userJSON             string
		expectCode           int
		expectBodyStartswith string
	}{
		{
			name:                 "Positive Case create user",
			path:                 "/users/create",
			userJSON:             "{\"name\":\"test create\",\"email\":\"testcreate@email.com\",\"password\":\"123\"}",
			expectCode:           http.StatusCreated,
			expectBodyStartswith: "{\"message\":\"success create new user\",\"user\":",
		},
		{
			name:                 "Negative Case create user with blank field",
			path:                 "/users/create",
			userJSON:             `{"name":"test create","password":"123"}`,
			expectCode:           http.StatusBadRequest,
			expectBodyStartswith: `{"messages":"email must be valid`,
		},
	}
	e := InitEchoTest()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(testCase.userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, CreateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartswith))
		}
	}
}

func TestUpdateUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userID               string
		userJSON             string
		expectCode           int
		expectBodyStartswith string
	}{
		{
			name:                 "Positive Case Edit user",
			path:                 "/users/:id",
			userID:               "1",
			userJSON:             `{"name":"test edit","email":"testedit@email.com","password":"123"}`,
			expectCode:           http.StatusOK,
			expectBodyStartswith: "{\"messages\":\"success update user\",\"users\":",
		},
		{
			name:                 "Negative Case Edit user with Invalid parameter",
			path:                 "/users/:id",
			userID:               "daa",
			userJSON:             `{"name":"test edit","email":"testedit@email.com","password":"123"}`,
			expectCode:           http.StatusBadRequest,
			expectBodyStartswith: "{\"messages\":\"Invalid ID parameter",
		},
	}
	e := InitEchoTest()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, testCase.path, strings.NewReader(testCase.userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, UpdateUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartswith))
		}
	}
}

func TestDeleteUserController(t *testing.T) {
	testCases := []struct {
		name                 string
		path                 string
		userID               string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "Positive Case Delete User By ID",
			path:                 "/users/:id",
			userID:               "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"id\":",
		},
		{
			name:                 "Negative Case Delete User Invalid ID Parameter",
			path:                 "/users/:id",
			userID:               "dsa",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"messages\":\"Invalid ID Parameter",
		},
	}

	e := InitEchoTest()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.userID)

		//Assertion
		if assert.NoError(t, DeleteUserController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			//assert.Equal(t, testCase.expectBodyStartsWith, body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
		}
	}
}

func TestLoginUserController(t *testing.T) {
	var testCases = []struct {
		name                 string
		path                 string
		userJSON             string
		expectCode           int
		expectBodyStartswith string
	}{
		{
			name:                 "Positive Case Login",
			path:                 "/login",
			userJSON:             `{"email":"alterra@academy.id","password":"123"}`,
			expectCode:           http.StatusOK,
			expectBodyStartswith: "{\"message\":\"success login user\",\"user\":",
		},
		{
			name:                 "Negative Case Login",
			path:                 "/login",
			userJSON:             `{"email":"notexists@email.com","password":"123"}`,
			expectCode:           http.StatusBadRequest,
			expectBodyStartswith: "{\"messages\":\"record not found",
		},
	}
	e := InitEchoTest()

	// create user for login
	if err := CreateUserMocking(); err != nil {
		panic(err)
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, testCase.path, strings.NewReader(testCase.userJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, LoginUserController(c)) {
			assert.Equal(t, testCase.expectCode, rec.Code)
			// assert.Equal(t, testCase.expectBodyStartswith, body)
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.expectBodyStartswith))
		}
	}
}
