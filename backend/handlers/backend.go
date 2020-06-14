package handlers

import (
	"github.com/gin-gonic/gin"
	"gonas/models"
	"gonas/utils"
	"net/http"
)

type userInfo struct {
	Roles [1]string
	Introduction string
	Avatar string
	Name string
}

func Login(c *gin.Context) (err error) {
	var user models.User
	err = c.BindJSON(&user)
	token, err := user.Login()
	if err != nil {
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeProcessFailed, "校验失败，请重试！")
		return
	}
	if token == "" {
		utils.ClientJson(c, http.StatusBadRequest, token, utils.CodeProcessFailed, "用户名密码不匹配！")
		return
	}
	utils.ClientJson(c, http.StatusOK, token, utils.CodeSuccess, "success")
	return
}

func Logout(c *gin.Context) error {
	utils.ClientJson(c, http.StatusOK, "", utils.CodeSuccess, "success")
	return nil
}




func AllUsers(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"user1": "user1"})
	return nil
}

func UserInfo(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": userInfo{
		Roles:        [1]string{0:"admin"},
		Introduction: "I am a super administrator",
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:         "Super Admin",
	}})
	return nil
}
