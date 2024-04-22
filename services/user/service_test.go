package user

import (
	"context"
	"testing"
	"users/domain"
	"users/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	user = domain.User{
		ID:        uuid.NewString(),
		Username:  "poloyacero1",
		Email:     "akosipoloyacero1@gmail.com",
		Birthdate: "1987-12-18",
	}
)

func TestDeleteUser(t *testing.T) {
	mockUserRepo := new(mocks.UserMockRepository)

	mockUserRepo.On("Delete").Return(nil)

	userTestService := NewService(mockUserRepo)

	err := userTestService.DeleteUser(context.TODO(), uuid.MustParse(user.ID))

	mockUserRepo.AssertExpectations(t)

	assert.Nil(t, err)
}

func TestUpdateUsersValid(t *testing.T) {
	mockUserRepo := new(mocks.UserMockRepository)

	mockUserRepo.On("Update").Return(&user, nil)

	userTestService := NewService(mockUserRepo)

	result, _ := userTestService.UpdateUsers(context.TODO(), user, uuid.MustParse(user.ID))

	mockUserRepo.AssertExpectations(t)

	userResponse := &domain.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      36,
		Base:     result.Base,
	}

	assert.Equal(t, userResponse, result)
}

func TestFindUserValid(t *testing.T) {
	mockUserRepo := new(mocks.UserMockRepository)

	mockUserRepo.On("Find").Return(&user, nil)

	userTestService := NewService(mockUserRepo)

	result, _ := userTestService.FindUser(context.TODO(), uuid.MustParse(user.ID))

	mockUserRepo.AssertExpectations(t)

	userResponse := &domain.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      36,
		Base:     result.Base,
	}

	assert.Equal(t, userResponse, result)
}

func TestCreateUsersValid(t *testing.T) {
	mockUserRepo := new(mocks.UserMockRepository)

	mockUserRepo.On("Save").Return(&user, nil)

	userTestService := NewService(mockUserRepo)

	result, _ := userTestService.CreateUsers(context.TODO(), user)

	mockUserRepo.AssertExpectations(t)

	userResponse := &domain.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      36,
		Base:     result.Base,
	}

	assert.Equal(t, userResponse, result)
}

func TestBirthdateValid(t *testing.T) {
	userTestService := NewService(nil)

	age := userTestService.ageParser("1987-12-18")

	assert.NotEmpty(t, age)

	assert.Equal(t, uint8(36), age)
}
