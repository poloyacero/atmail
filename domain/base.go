package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        uuid.UUID      `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (b *Base) BeforeCreate(_ *gorm.DB) error {
	if b.ID != uuid.Nil {
		return nil
	}
	b.ID = uuid.New()
	return nil
}

type Data struct {
	Data map[string]string
}
