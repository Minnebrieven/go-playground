package repository

import (
	"mvc-to-ca/model"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *UserRepositoryMock {
	return &UserRepositoryMock{}
}

func (u *UserRepositoryMock) GetAllUsers() ([]model.User, error) {
	args := u.Called()
	return args.Get(0).([]model.User), args.Error(1)
}

func (u *UserRepositoryMock) GetUser(userData model.User) (model.User, error) {
	args := u.Called(userData)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepositoryMock) Login(userData model.User) (model.User, error) {
	args := u.Called(userData)
	return args.Get(0).(model.User), args.Error(1)
}

func (u *UserRepositoryMock) CreateUser(userData model.User) error {
	args := u.Called(userData)
	return args.Error(0)
}
