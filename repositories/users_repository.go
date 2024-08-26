package repositories

import (
	"github.com/duylamasd/gorm-try/interfaces"
)

type UsersRepository struct {
	db interfaces.Database
}

func NewUsersRepository(db interfaces.Database) interfaces.UsersRepository {
	return &UsersRepository{db: db}
}

func (repository *UsersRepository) FindById() string {
	return "User"
}
