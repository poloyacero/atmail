package user

import (
	"context"
	"time"
	"users/domain"

	"github.com/google/uuid"
)

type Service struct {
	userRepository domain.UserRepository
}

func (s *Service) CreateUsers(ctx context.Context, request domain.User) (*domain.UserResponse, error) {
	result, err := s.userRepository.Save(ctx, &request)
	if err != nil {
		return nil, err
	}

	userResponse := &domain.UserResponse{}

	userResponse.ID = result.ID
	userResponse.Email = result.Email
	userResponse.Username = result.Username
	userResponse.Age = s.ageParser(result.Birthdate)
	userResponse.Base = result.Base

	return userResponse, nil
}

func (s *Service) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	err := s.userRepository.Delete(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) FindUser(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	result, err := s.userRepository.Find(ctx, userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) UpdateUsers(ctx context.Context, request domain.User, userID uuid.UUID) (*domain.UserResponse, error) {
	user := &domain.User{
		ID:        userID.String(),
		Birthdate: request.Birthdate,
	}

	result, err := s.userRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	userResponse := &domain.UserResponse{}

	userResponse.ID = result.ID
	userResponse.Email = result.Email
	userResponse.Username = result.Username
	userResponse.Age = s.ageParser(result.Birthdate)
	userResponse.Base = result.Base

	return userResponse, nil
}

func (s *Service) ageParser(birthdate string) uint8 {
	layout := "2006-01-02T00:00:00Z"
	date, _ := time.Parse(layout, birthdate)
	now := time.Now()
	age := now.Sub(date).Hours() / 24 / 365

	return uint8(age)
}

func NewService(userRepository domain.UserRepository) Service {
	return Service{userRepository}
}
