package handlers

import (
	"github.com/gin-gonic/gin"
	"gonas/models"
	"gonas/utils"
)

func AllDirectories(c *gin.Context) error {
	utils.JsonRepose(c)
	return nil
}

func CreateDirectories(c *gin.Context) (err error) {
	var directory models.Directory
	err = c.BindJSON(&directory)
	if err != nil {
		utils.JsonRepose(c, "", "创建文件夹失败！", utils.CodeProcessFailed)
		return
	}
	userID, err := models.CheckAuth(c)
	if err != nil {
		utils.JsonRepose(c, "", "创建文件夹失败！", utils.CodeProcessFailed)
		return
	}
	directory.UserID = userID
	if ok, err := directory.Create(); !ok{
		utils.JsonRepose(c, "", "创建文件夹失败！", utils.CodeProcessFailed)
		return err
	}
	utils.JsonRepose(c, directory)
	return
}
