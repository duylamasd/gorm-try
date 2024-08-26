package services

import "github.com/duylamasd/gorm-try/interfaces"

type UsersService struct {
	usersRepository interfaces.UsersRepository
}

func NewUsersService(usersRepository interfaces.UsersRepository) interfaces.UsersService {
	return &UsersService{usersRepository: usersRepository}
}

func (service *UsersService) FindById() string {
	return service.usersRepository.FindById()
}
