package router

import (
	"auth-server/handlers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	{
		r.POST("/test/auth", handlers.TestAuthorization)
	}
}
