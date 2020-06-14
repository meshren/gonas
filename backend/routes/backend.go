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
	backend.POST("/login", ErrWrapper(handlers.Login))
	backend.Use(middlware.Auth())
	{
		backend.GET("/user/info", ErrWrapper(handlers.UserInfo))
		backend.POST("/user/logout", ErrWrapper(handlers.Logout))
		backend.GET("/files", ErrWrapper(handlers.AllFiles))
		backend.POST("/files", ErrWrapper(handlers.UploadFiles)) // 上传文件
		backend.GET("/users", ErrWrapper(handlers.AllUsers))
		backend.POST("/directories", ErrWrapper(handlers.CreateDirectories))
		backend.GET("/directories", ErrWrapper(handlers.AllDirectories))
	}
}
