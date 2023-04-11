package database

import (
	"crud-echo-gorm/config"
	"crud-echo-gorm/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func GetBook(bookID int) (interface{}, error) {
	var book models.Book
	book.ID = uint(bookID)

	if err := config.DB.First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func CreateBook(b models.Book) (interface{}, error) {
	err := config.DB.Create(&b).Error

	if err != nil {
		return nil, err
	}

	return b, nil
}

func UpdateBook(bookID uint, b models.Book) (interface{}, error) {
	book := models.Book{}
	book.ID = bookID
	if err := config.DB.First(&book).Error; err != nil {
		return nil, err
	}

	book.Title = b.Title
	book.Author = b.Author
	book.Publisher = b.Publisher

	if err := config.DB.Save(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func DeleteBook(bookID int) (interface{}, error) {
	err := config.DB.Delete(&models.Book{}, bookID).Error

	if err != nil {
		return nil, err
	}
	return bookID, nil
}
