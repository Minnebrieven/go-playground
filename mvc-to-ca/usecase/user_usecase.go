package usecase

import (
	"mvc-to-ca/middleware"
	"mvc-to-ca/model"
	"mvc-to-ca/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]model.User, error)
	GetUser(model.User) (model.User, error)
	Login(model.User) (string, error)
	CreateUser(model.User) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepo}
}

func (us *userUsecase) GetAllUsers() ([]model.User, error) {
	user, err := us.userRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userUsecase) GetUser(user model.User) (model.User, error) {
	user, err := us.userRepository.GetUser(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (us *userUsecase) Login(user model.User) (string, error) {
	user, err := us.userRepository.Login(user)
	if err != nil {
		return "", err
	}

	token, err := middleware.CreateToken(int(user.ID), user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (us *userUsecase) CreateUser(dataUser model.User) error {
	if err := us.userRepository.CreateUser(dataUser); err != nil {
		return err
	}
	return nil
}
