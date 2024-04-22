package mocks

import (
	"context"
	"users/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserMockService struct {
	mock.Mock
}

func (m *UserMockService) Create(ctx context.Context, request domain.User) (*domain.UserResponse, error) {
	args := m.Called(ctx, request)

	result := args.Get(0)
	return result.(*domain.UserResponse), args.Error(1)
}

func (m *UserMockService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	args := m.Called(ctx, userID)

	return args.Error(1)
}

func (m *UserMockService) FindUser(ctx context.Context, userID uuid.UUID) (*domain.UserResponse, error) {
	args := m.Called(ctx, userID)

	result := args.Get(0)
	return result.(*domain.UserResponse), args.Error(1)
}

func (m *UserMockService) UpdateUsers(ctx context.Context, request domain.User, userID uuid.UUID) (*domain.UserResponse, error) {
	args := m.Called(ctx, request, userID)

	result := args.Get(0)
	return result.(*domain.UserResponse), args.Error(1)
}
