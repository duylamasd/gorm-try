package routes

import (
	"github.com/duylamasd/gorm-try/interfaces"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func Init(engine *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func(usersControllers interfaces.UsersController) error {
		engine.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})

		usersGroup := engine.Group("/users")
		{
			usersGroup.GET("/", usersControllers.FindById)
		}

		return nil
	})

	return err
}
