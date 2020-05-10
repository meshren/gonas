package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gonas/handlers"
	"gonas/middlware"
)

func Run() {
	router := gin.Default()
	router.Static("/admin", "./html")
	router.Static("/static", "./html/static")
	router.GET("/test", handlers.TestModel)
	// 后台文件管理系统
	backend := router.Group("/backend")
	backend.POST("/login", handlers.Login)
	backend.Use(middlware.Auth())
	{
		backend.GET("/files", handlers.AllFiles)
		backend.GET("/users", handlers.AllUsers)
		backend.GET("/user/info", handlers.UserInfo)
		backend.POST("/user/logout", handlers.Logout)
	}

	// 客户端通用接口
	api := router.Group("/api")
	{
		api.GET("/login", handlers.Login)
	}

	port := viper.GetString("backend.golang.port")

	router.Run(":" + port)
}
