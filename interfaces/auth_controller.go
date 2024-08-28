package interfaces

import "github.com/gin-gonic/gin"

type AuthController interface {
	SignUp(ctx *gin.Context)
}
