package routes

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gonas/handlers"
	"log"
	"net/http"
)

type HandlerFunc func(c *gin.Context) (err error)

// 统一处理error
func ErrWrapper(handler HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("check error")
		err := handler(c)
		log.Println(err)
	}
}

func user(c *gin.Context) (err error) {
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
	return errors.New("Are you ok? ")
}

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

