package interfaces

import (
	"github.com/duylamasd/gorm-try/models"
	"github.com/duylamasd/gorm-try/schemas"
)

type UsersRepository interface {
	FindById() string
	Create(user *schemas.SignupBody) (*models.User, error)
}
