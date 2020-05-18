package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gonas/handlers"
)

func Run() {
	router := gin.Default()
	router.GET("/test", handlers.TestModel)
	// 用户后台文件管理系统
	backend(router)
	// 客户端通用接口
	api(router)

	port := viper.GetString("backend.golang.port")

	router.Run(":" + port)
}
