package handler

import (
	"github.com/gin-gonic/gin"
)

func UpdateTaskHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "UPDATE task",
	})
}
