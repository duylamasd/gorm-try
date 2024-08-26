package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func Init(engine *gin.Engine, container *dig.Container) error {
	err := container.Invoke(func() error {
		engine.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!",
			})
		})

		return nil
	})

	return err
}
