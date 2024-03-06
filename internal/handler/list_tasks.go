package handler

import (
	"github.com/gin-gonic/gin"
)

func ListTaskHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "GET all task",
	})
}
