package router

import (
	"go-app/pkg/handler"

	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	// Route group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", handler.Login)
		v1.POST("/submit", handler.Submit)
		v1.POST("/read", handler.Read)
	}

	// Route group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", handler.Login)
		v2.POST("/submit", handler.Submit)
		v2.POST("/read", handler.Read)
	}

	router.Run(":8080")

	return router
}
