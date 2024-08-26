package interfaces

import "github.com/gin-gonic/gin"

type UsersController interface {
	FindById(ctx *gin.Context)
}
