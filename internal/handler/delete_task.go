package handler

import (
	"github.com/gin-gonic/gin"
)

func DeleteTaskHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "DELETE task",
	})
}
