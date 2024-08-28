package services

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUsersService)
	_ = container.Provide(NewAuthService)

	return nil
}
