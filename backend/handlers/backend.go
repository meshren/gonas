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

func Login(c *gin.Context) {
	// 1. 验证表单
	var manager models.Manager
	err := c.BindJSON(&manager)
	if err != nil {
		utils.ErrDetail(err)
		ClientJson(c, http.StatusOK, "", CodeInvalidParam, "参数不正确！")
		return
	}
	if len(manager.Username) < 6 || len(manager.Password) < 6 {
		ClientJson(c, http.StatusOK, "", CodeInvalidParam, "用户名或密码无效！")
		return
	}
	// 2. 验证密码
	if ok, _ := manager.CheckPassword(); !ok {
		ClientJson(c, http.StatusOK, "", CodeInvalidParam, "用户名密码不匹配！")
		return
	}
	// 3. 生成token
	token, err := models.GenToken(&manager)
	if err != nil {
		utils.ErrDetail(err)
		ClientJson(c, http.StatusOK, "", CodeInvalidParam, "生成用户token出错！")
		return
	}
	tokenEncrypt, err := models.EncryptToken(&token)
	if err != nil {
		utils.ErrDetail(err)
		ClientJson(c, http.StatusOK, "", CodeInvalidParam, "加密用户token出错！")
		return
	}
	ClientJson(c, http.StatusOK, tokenEncrypt, CodeSuccess, "success")
}

func Logout(c *gin.Context)  {
	ClientJson(c, http.StatusOK, "", CodeSuccess, "success")
}

func AllFiles(c *gin.Context) {
	var file models.File
	files, err := file.GetAll("created_at DESC", 10, 0)
	if err != nil {
		utils.ErrDetail(err)
		ClientJson(c, http.StatusBadRequest, files, CodeProcessFailed, "查询失败！")
	}
	ClientJson(c, http.StatusOK, files, CodeSuccess, "success")
}

func AllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"user1": "user1"})
}

func UserInfo(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"code": 20000, "data": userInfo{
		Roles:        [1]string{0:"admin"},
		Introduction: "I am a super administrator",
		Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		Name:         "Super Admin",
	}})
}
