package database

import (
	"crud-echo-gorm/config"
	"crud-echo-gorm/models"
)

func GetBlogs() (interface{}, error) {
	var blogs []models.Blog

	if err := config.DB.Joins("User").Find(&blogs).Error; err != nil {
		return nil, err
	}
	return blogs, nil
}

func GetBlog(blogID int) (interface{}, error) {
	var blog models.Blog
	blog.ID = uint(blogID)

	if err := config.DB.Joins("User").Find(&blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

func CreateBlog(b models.Blog) (interface{}, error) {
	if err := config.DB.Create(&b).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Joins("User").Find(&b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

func UpdateBlog(blogID uint, b models.Blog) (interface{}, error) {
	blog := models.Blog{}
	blog.ID = blogID
	if err := config.DB.Joins("User").Find(&blog).Error; err != nil {
		return nil, err
	}

	blog.Title = b.Title
	blog.Content = b.Content
	blog.UserID = b.UserID

	if err := config.DB.Save(&blog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}

func DeleteBlog(blogID int) (interface{}, error) {
	err := config.DB.Delete(&models.Blog{}, blogID).Error

	if err != nil {
		return nil, err
	}
	return blogID, nil
}
