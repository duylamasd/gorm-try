package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	UID         uuid.UUID `gorm:"not null"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	TTL         time.Time
	CompletedAt time.Time
}

func (Todo) TableName() string {
	return "todos"
}
