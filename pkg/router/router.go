package router

import (
	"go-app/pkg/handler"

	"github.com/BoyYangZai/go-server-lib/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.POST("/verify-code", handler.VerifyCode)
		user.POST("/registry", handler.Registry)
		user.POST("/login", handler.Login)
	}

	auth_test := router.Group("/auth-test")
	{
		auth_test.Use(jwt.AuthMiddleware())
		auth_test.GET("/", handler.Submit)
	}

	router.Run(":8080")

	return router
}
