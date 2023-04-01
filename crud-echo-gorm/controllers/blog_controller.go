package controllers

import (
	"net/http"
	"strconv"

	"crud-echo-gorm/lib/database"
	"crud-echo-gorm/models"

	"github.com/labstack/echo/v4"
)

// get all blogs
func GetBlogsController(c echo.Context) error {
	blogs, err := database.GetBlogs()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all blogs",
		"blogs":    blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	blog, err := database.GetBlog(blogID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get blog",
		"blogs":    blog,
	})
}

// create new blog
func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	result, err := database.CreateBlog(blog)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new blog",
		"blog":    result,
	})
}

// update blog by id
func UpdateBlogController(c echo.Context) error {

	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	blog := models.Blog{}
	c.Bind(&blog)

	result, err := database.UpdateBlog(uint(blogID), blog)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update blog",
		"blogs":    result,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	blogID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := database.DeleteBlog(blogID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete blog",
		"id":       result,
	})
}
