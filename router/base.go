package router

import (
	_ "auth-server/docs"
	"auth-server/handlers"
	"auth-server/middleware"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/user/register", handlers.Register)
	r.POST("/user/login", handlers.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthorizationMiddleware())
	UserRouter(api)

	return r
}
