package mysql

import (
	"context"
	"fmt"
	"users/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	db := u.db.WithContext(ctx)

	if err := db.Create(user).Error; err != nil {
		return nil, fmt.Errorf("userRepository: %w", err)
	}

	return user, nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	db := u.db.WithContext(ctx)
	modelSource := &domain.User{}

	if err := db.Model(modelSource).Clauses(clause.Returning{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return nil, fmt.Errorf("userRepository: %w", err)
	}

	modelSource, err := u.Find(ctx, uuid.MustParse(user.ID))
	if err != nil {
		return nil, err
	}

	return modelSource, nil
}

func (u *userRepository) Find(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	db := u.db.WithContext(ctx)
	modelSource := &domain.User{}

	err := db.Where("id = ?", userID).First(modelSource).Error
	if err != nil {
		return nil, err
	}

	return modelSource, nil
}

func (u *userRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	user, err := u.Find(ctx, userID)
	if err != nil {
		return err
	}

	if err := u.db.Delete(&user).Error; err != nil {
		return fmt.Errorf("userRepository: %w", err)
	}

	return nil
}

func (u *userRepository) IsExist(ctx context.Context, dataType, data string) bool {
	modelSource := &domain.User{}
	if err := u.db.WithContext(ctx).Where(dataType+" = ?", data).First(modelSource).Error; err != nil {
		return true
	}

	return false
}

func NewSourceRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}
