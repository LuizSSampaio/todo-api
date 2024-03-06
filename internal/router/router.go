package router

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()

	initializeRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
