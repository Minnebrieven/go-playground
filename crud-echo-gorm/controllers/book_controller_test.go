package controllers

import (
	"crud-echo-gorm/config"
	"crud-echo-gorm/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BookTestCase struct {
	Name                 string
	Path                 string
	BookID               string
	BookJSON             string
	ExpectCode           int
	ExpectBodyStartsWith string
}

func init() {
	config.InitDBTest()
	if err := CreateBookMocking(); err != nil {
		panic(err)
	}
}

func CreateBookMocking() error {
	book := models.Book{
		Title:     "The Great Gatsby",
		Author:    "F. Scott Fitzgerald",
		Publisher: "Simon & Schuster",
	}

	if err := config.DB.Save(&book).Error; err != nil {
		return err
	}

	return nil
}

func TestGetBooksController(t *testing.T) {
	testCases := []BookTestCase{
		{
			Name:                 "Positive Case Get All Books",
			Path:                 "/books",
			ExpectCode:           http.StatusOK,
			ExpectBodyStartsWith: "{\"books\":[",
		},
		// {
		// 	Name:                 "Negative Case wrong path",
		// 	Path:                 "/book",
		// 	ExpectCode:           http.StatusNotFound,
		// 	ExpectBodyStartsWith: "{\"message\":\"Not Found",
		// },
	}

	e := InitEchoTest()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)

		//Assertion
		if assert.NoError(t, GetBooksController(c)) {
			assert.Equal(t, testCase.ExpectCode, rec.Code)
			body := rec.Body.String()
			// assert.Equal(t, testCase.ExpectBodyStartsWith, body)
			assert.True(t, strings.HasPrefix(body, testCase.ExpectBodyStartsWith))
		}
	}
}

func TestGetBookController(t *testing.T) {
	testCases := []BookTestCase{
		{
			Name:                 "Positive Case Get Books By ID",
			Path:                 "/books/:id",
			BookID:               "1",
			ExpectCode:           http.StatusOK,
			ExpectBodyStartsWith: "{\"books\":",
		},
		{
			Name:                 "Negative Case Get Books By ID not Exists",
			Path:                 "/books/:id",
			BookID:               "9999",
			ExpectCode:           http.StatusBadRequest,
			ExpectBodyStartsWith: "{\"messages\":\"record not found",
		},
	}

	if err := CreateBookMocking(); err != nil {
		panic(err)
	}

	e := InitEchoTest()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.BookID)

		//Assertion
		if assert.NoError(t, GetBookController(c)) {
			assert.Equal(t, testCase.ExpectCode, rec.Code)
			body := rec.Body.String()
			// assert.Equal(t, testCase.ExpectBodyStartsWith, body)
			assert.True(t, strings.HasPrefix(body, testCase.ExpectBodyStartsWith))
		}
	}
}

func TestCreateBookController(t *testing.T) {
	testCases := []BookTestCase{
		{
			Name:                 "Positive Case Create Book",
			Path:                 "/books",
			BookJSON:             `{"title":"test title","author":"test author","publisher":"test publisher"}`,
			ExpectCode:           http.StatusCreated,
			ExpectBodyStartsWith: "{\"book\":",
		},
		{
			Name:                 "Negative Case create user with blank field",
			Path:                 "/books",
			BookJSON:             `"author":"test author","publisher":"test publisher"}`,
			ExpectCode:           http.StatusBadRequest,
			ExpectBodyStartsWith: `{"messages":"title must be valid`,
		},
	}

	e := InitEchoTest()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, testCase.Path, strings.NewReader(testCase.BookJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, CreateBookController(c)) {
			assert.Equal(t, testCase.ExpectCode, rec.Code)
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.ExpectBodyStartsWith))
		}
	}
}

func TestUpdateBookController(t *testing.T) {
	testCases := []BookTestCase{
		{
			Name:                 "Positive Case Edit Book",
			Path:                 "/books/:id",
			BookID:               "1",
			BookJSON:             `{"title":"Killing The Mocking Bird","author":"Harper Lee","publisher":"J. B. Lippincott & Co."}`,
			ExpectCode:           http.StatusOK,
			ExpectBodyStartsWith: "{\"books\":",
		},
		{
			Name:                 "Negative Case Edit Book Invalid ID Parameter",
			Path:                 "/books/:id",
			BookID:               "dsdasd",
			BookJSON:             `{"title":"Killing The Mocking Bird","author":"Harper Lee","publisher":"J. B. Lippincott & Co."}`,
			ExpectCode:           http.StatusBadRequest,
			ExpectBodyStartsWith: "{\"messages\":\"Invalid ID Parameter",
		},
	}

	e := InitEchoTest()
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(testCase.BookJSON))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.BookID)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, UpdateBookController(c)) {
			assert.Equal(t, testCase.ExpectCode, rec.Code)
			//assert.Equal(t, testCase.ExpectBodyStartsWith, rec.Body.String())
			assert.True(t, strings.HasPrefix(rec.Body.String(), testCase.ExpectBodyStartsWith))
		}
	}
}

func TestDeleteBookController(t *testing.T) {
	testCases := []BookTestCase{
		{
			Name:                 "Positive Case Delete Book",
			Path:                 "/books/:id",
			BookID:               "1",
			ExpectCode:           http.StatusOK,
			ExpectBodyStartsWith: "{\"id\":",
		},
		{
			Name:                 "Negative Case Delete Book Invalid ID Parameter",
			Path:                 "/books/:id",
			BookID:               "afsdas",
			ExpectCode:           http.StatusBadRequest,
			ExpectBodyStartsWith: "{\"messages\":\"Invalid ID Parameter",
		},
	}

	e := InitEchoTest()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	for _, testCase := range testCases {
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.Path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.BookID)

		//Assertion
		if assert.NoError(t, DeleteBookController(c)) {
			assert.Equal(t, testCase.ExpectCode, rec.Code)
			body := rec.Body.String()
			//assert.Equal(t, testCase.expectBodyStartsWith, body)
			assert.True(t, strings.HasPrefix(body, testCase.ExpectBodyStartsWith))
		}
	}
}
