package database

import (
	"crud-echo-gorm/config"
	"crud-echo-gorm/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(userID uint) (interface{}, error) {
	var user models.User
	user.ID = userID

	if err := config.DB.First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(u models.User) (interface{}, error) {
	err := config.DB.Create(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func UpdateUser(userID uint, u models.User) (interface{}, error) {
	user := models.User{}
	user.ID = userID
	config.DB.First(&user)

	user.Name = u.Name
	user.Email = u.Email
	user.Password = u.Password

	err := config.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(userID int) (interface{}, error) {
	err := config.DB.Delete(&models.User{}, userID).Error

	if err != nil {
		return nil, err
	}
	return userID, nil
}
