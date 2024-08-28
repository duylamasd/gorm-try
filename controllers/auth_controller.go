package controllers

import (
	"net/http"

	"github.com/duylamasd/gorm-try/interfaces"
	"github.com/duylamasd/gorm-try/schemas"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service interfaces.AuthService
}

func NewAuthController(service interfaces.AuthService) interfaces.AuthController {
	return &AuthController{service: service}
}

func (controller *AuthController) SignUp(ctx *gin.Context) {
	var body schemas.SignupBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenInfo, err := controller.service.SignUp(ctx, &body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tokenInfo)
}
