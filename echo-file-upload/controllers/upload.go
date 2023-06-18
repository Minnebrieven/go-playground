package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	htmlFile := `<!doctype html>
	<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Single file upload</title>
	</head>
	<body>
	<h1>Upload single file with fields</h1>
	
	<form action="/upload" method="post" enctype="multipart/form-data">
		Files: <input type="file" name="file"><br><br>
		<input type="submit" value="Submit">
	</form>
	</body>
	</html>
	`
	return c.HTML(200, htmlFile)
}

func Upload(c echo.Context) error {
	//-----------
	// Read file
	//-----------

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	fileExtension := path.Ext(file.Filename)
	src, err := file.Open()
	if err != nil {
		return err
	}
	
	defer src.Close()

	destinationFileRoute := fmt.Sprintf("assets/%s", file.Filename)

	// Destination
	dst, err := os.Create(destinationFileRoute)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully. File extension : %s</p>", file.Filename, fileExtension))
}
