package usecase

import (
	"mvc-to-ca/model"
	"mvc-to-ca/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	mockUsers := []model.User{
		{
			ID:        1,
			Email:     "mockEmail@email.com",
			Password:  "123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockUserRepo := repository.NewUserRepositoryMock()
	mockUserRepo.On("GetAllUsers").Return(mockUsers, nil)

	getAllService := &repository.UserRepositoryMock{Mock: mockUserRepo.Mock}
	users, err := getAllService.GetAllUsers()

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, mockUsers[0].Email, users[0].Email)
	assert.Equal(t, mockUsers[0].Password, users[0].Password)

}

func TestGetUser(t *testing.T) {
	mockUser := model.User{
		ID:        1,
		Email:     "mockEmail@email.com",
		Password:  "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUserRepo := repository.NewUserRepositoryMock()
	mockUserRepo.On("GetUser", mockUser).Return(mockUser, nil)

	getService := &repository.UserRepositoryMock{Mock: mockUserRepo.Mock}
	users, err := getService.GetUser(mockUser)

	assert.Nil(t, err)
	assert.NotNil(t, users)

	assert.Equal(t, mockUser.Email, users.Email)
	assert.Equal(t, mockUser.Password, users.Password)
}

// func TestLogin(t *testing.T) {
// 	mockToken := ""

// 	mockUserRepo := repository.NewUserRepositoryMock()
// 	mockUserRepo.On("GetUser", mockUser).Return(mockUser, nil)

// 	getService := &repository.UserRepositoryMock{Mock: mockUserRepo.Mock}
// 	users, err := getService.GetUser(mockUser)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, users)

// 	assert.Equal(t, mockUser.Email, users.Email)
// 	assert.Equal(t, mockUser.Password, users.Password)
// }

// func TestCreateUser(t *testing.T) {
// 	mockUser := model.User{
// 		Email:    "testcreate@email.com",
// 		Password: "123",
// 	}

// 	userRMock.Mock.On("CreateRepository", user).Return(user, nil)
// 	users, err := userSMock.CreateService(user)

// 	assert.Nil(t, err)
// 	assert.NotNil(t, users)

// 	assert.Equal(t, user.Email, users.Email)
// 	assert.Equal(t, user.Password, users.Password)
// }
