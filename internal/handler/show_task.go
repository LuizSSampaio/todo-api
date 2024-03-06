package handler

import (
	"github.com/gin-gonic/gin"
)

func ShowTaskHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "GET task",
	})
}
