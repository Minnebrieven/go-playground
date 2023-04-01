package controllers

import (
	"net/http"
	"strconv"

	"crud-echo-gorm/lib/database"
	"crud-echo-gorm/models"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all books",
		"books":    books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book, err := database.GetBook(bookID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get book",
		"books":    book,
	})
}

// create new book
func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	result, err := database.CreateBook(book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new book",
		"book":    result,
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {

	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book := models.Book{}
	c.Bind(&book)

	result, err := database.UpdateBook(uint(bookID), book)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update book",
		"books":    result,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := database.DeleteBook(bookID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete book",
		"id":       result,
	})
}
