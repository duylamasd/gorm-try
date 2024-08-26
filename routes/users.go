package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersResource struct{}

func (rs UsersResource) Routes(rg *gin.RouterGroup) {
	users := rg.Group("/users")
	{
		users.GET("/", rs.get)
	}
}

func (rs UsersResource) get(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "users")
}
