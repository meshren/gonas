package routes

import (
	"github.com/gin-gonic/gin"
	"gonas/handlers"
)

func api(router *gin.Engine)  {
	api := router.Group("/api")
	api.GET("/login", handlers.Login)
}
