package router

import (
	"github.com/LuizSSampaio/todo-api/internal/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1/")
	{
		v1.GET("task", handler.ShowTaskHandler)
		v1.POST("task", handler.CreateTaskHandler)
		v1.DELETE("task", handler.DeleteTaskHandler)
		v1.PUT("task", handler.UpdateTaskHandler)
		v1.GET("tasks", handler.ListTaskHandler)
	}
}
