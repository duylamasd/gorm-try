package controllers

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	_ = container.Provide(NewUsersController)
	_ = container.Provide(NewAuthController)

	return nil
}
