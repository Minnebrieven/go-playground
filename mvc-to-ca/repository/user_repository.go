package repository

import (
	"mvc-to-ca/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	GetUser(model.User) (model.User, error)
	Login(model.User) (model.User, error)
	CreateUser(model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetAllUsers() ([]model.User, error) {
	users := []model.User{}
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *userRepository) GetUser(user model.User) (model.User, error) {
	err := ur.db.First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) Login(user model.User) (model.User, error) {
	err := ur.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *userRepository) CreateUser(userData model.User) error {
	return ur.db.Create(&userData).Error
}
