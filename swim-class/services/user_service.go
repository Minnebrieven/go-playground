package services

import (
	"swim-class/middlewares"
	"swim-class/models"
	"swim-class/repositories"
)

type UserService interface {
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) *userService {
	return &userService{userRepository: userRepo}
}

func (us *userService) GetAllUsersService() ([]models.User, error) {
	user, err := us.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userService) GetUserService(userData models.User) (models.User, error) {
	user, err := us.userRepository.GetUser(userData)
	return user, err
}

func (us *userService) CreateUserService(dataUser models.User) error {
	if err := us.userRepository.CreateUser(dataUser); err != nil {
		return err
	}
	return nil
}

func (us *userService) EditUserService(userData models.User) (*models.User, error) {
	//find record first if not exists return error
	_, err := us.userRepository.GetUser(userData)
	if err != nil {
		return nil, err
	}

	result, err := us.userRepository.UpdateUser(userData)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (us *userService) Login(userData models.User) (string, error) {
	user, err := us.userRepository.Login(userData)
	if err != nil {
		return "", err
	}

	token, err := middlewares.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}
