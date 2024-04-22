package domain

import (
	"context"

	"github.com/google/uuid"
)

type Role string

const (
	Customer   Role = "customer"
	Admin      Role = "admin"
	SuperAdmin Role = "super_admin"
)

func (r Role) String() string {
	if r == Customer || r == Admin || r == SuperAdmin {
		return string(r)
	}
	return string(Customer)
}

type UpdateUserRequest struct {
	Birthdate string `json:"birthdate" binding:"validDate"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username" binding:"required,unique"`
	Email    string `json:"email" binding:"required,email,unique"`
	Age      uint8  `json:"age,omitempty"`
	Base
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username" binding:"required,unique"`
	Email     string `json:"email" binding:"required,email,unique"`
	Birthdate string `json:"birthdate" binding:"required,validDate"`
	Base
}

type UserSuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

type UserRepository interface {
	Save(ctx context.Context, model *User) (*User, error)
	Update(ctx context.Context, model *User) (*User, error)
	Find(ctx context.Context, userID uuid.UUID) (*User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	IsExist(ctx context.Context, dataType, data string) bool
}
