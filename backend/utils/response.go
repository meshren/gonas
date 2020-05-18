package utils

import "github.com/gin-gonic/gin"

const (
	CodeSuccess       = 20000
	CodeInvalidParam  = 20001
	CodeProcessFailed = 40000
)

func ClientJson(c *gin.Context, httpStatus int, data interface{}, Code int, message string) {
	c.JSON(httpStatus, gin.H{"code": Code, "data": data, "message": message})
}
