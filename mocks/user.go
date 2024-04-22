package mocks

import (
	"context"
	"users/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type UserMockRepository struct {
	mock.Mock
}

func (mock *UserMockRepository) Save(ctx context.Context, model *domain.User) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

func (mock *UserMockRepository) Update(ctx context.Context, model *domain.User) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

func (mock *UserMockRepository) Find(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.User), args.Error(1)
}

func (mock *UserMockRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *UserMockRepository) IsExist(ctx context.Context, dataType, data string) bool {
	args := mock.Called()
	result := args.Get(0)
	return result.(bool)
}
