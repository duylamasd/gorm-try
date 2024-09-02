package repositories

import (
	"time"

	"github.com/duylamasd/gorm-try/interfaces"
	"github.com/duylamasd/gorm-try/models"
	"github.com/duylamasd/gorm-try/schemas"
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

func (repository *UsersRepository) Create(payload *schemas.SignupBody) (*models.User, error) {
	dateOfBirth, err := time.Parse("2006-01-02", payload.DateOfBirth)
	if err != nil {
		return nil, err
	}
	user := &models.User{Name: payload.Name, Email: payload.Email, Password: payload.Password, DateOfBirth: &dateOfBirth, FirstName: payload.FirstName, LastName: payload.LastName}

	res := repository.db.GetInstance().Create(user)

	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}
