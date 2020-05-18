package routes

import (
	"github.com/gin-gonic/gin"
	"gonas/handlers"
	"gonas/middlware"
)

func backend(router *gin.Engine) {
	router.Static("/admin", "./html")
	router.Static("/static", "./html/static")

	backend := router.Group("/backend")
	backend.POST("/login", handlers.Login)
	backend.Use(middlware.Auth())
	{
		backend.GET("/user/info", handlers.UserInfo)
		backend.POST("/user/logout", handlers.Logout)
		backend.GET("/files", handlers.AllFiles)
		backend.POST("/files", handlers.UploadFiles) // 上传文件
		backend.GET("/users", handlers.AllUsers)
		backend.POST("/folders", handlers.CreateFolder)
	}
}
