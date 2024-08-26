package controllers

import (
	"net/http"

	"github.com/duylamasd/gorm-try/interfaces"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	service interfaces.UsersService
}

func NewUsersController(service interfaces.UsersService) interfaces.UsersController {
	return &UsersController{service: service}
}

func (controller *UsersController) FindById(ctx *gin.Context) {
	res := controller.service.FindById()
	ctx.JSON(http.StatusOK, gin.H{"message": res})
}
