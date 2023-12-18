package handler

import (
	"go-app/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Submit(c *gin.Context) {
	// Implementation for submit

	c.JSON(http.StatusOK, gin.H{
		"msg":             "submit success",
		"currentAuthUser": service.GetAuthUser(),
	})

}

func Read(c *gin.Context) {
	// Implementation for read
}
