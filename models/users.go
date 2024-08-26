package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	Name        string     `gorm:"not null;index:,unique,expression:LOWER(name)"`
	Email       string     `gorm:"not null;index:,unique,expression:LOWER(name)"`
	Password    string     `gorm:"not null"`
	DateOfBirth *time.Time `gorm:"not null"`
	FirstName   string     `gorm:"not null"`
	LastName    string     `gorm:"not null"`
	Todos       []Todo     `gorm:"foreignKey:UID"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password == "" {
		err = errors.New("Password is required")
		return
	}

	hash, error := hashPassword(u.Password)
	if error != nil {
		err = error
		return
	}

	tx.Statement.SetColumn("Password", hash)
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if tx.Statement.Changed("Password") {
		if u.Password == "" {
			err = errors.New("Password is required")
			return
		}

		hash, error := hashPassword(u.Password)
		if error != nil {
			err = error
			return
		}
		tx.Statement.SetColumn("Password", hash)
	}

	return
}

func hashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	return string(bytes), err
}
