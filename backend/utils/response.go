package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CodeSuccess       = 20000
	CodeInvalidParam  = 20001
	CodeProcessFailed = 40000
)

type response struct {
	httpStatus int
	code int
	data interface{}
	message string
}

func ClientJson(c *gin.Context, httpStatus int, data interface{}, Code int, message string) {
	c.JSON(httpStatus, gin.H{"code": Code, "data": data, "message": message})
	c.Abort()
}

// 0data, 1message, 2response code, 3http status
func JsonRepose(c *gin.Context, items ...interface{})  {
	var d interface{}
	var r response
	r.httpStatus = http.StatusOK
	r.code = CodeSuccess
	r.data = d
	r.message = "ok"
	for i, v := range items{
		switch i {
		case 0:
			r.data = v
		case 1:
			if _, ok := v.(string); ok {
				r.message = v.(string)
			}
		case 2:
			if _, ok := v.(int); ok {
				r.httpStatus = v.(int)
			}
		case 3:
			if _, ok := v.(int); ok {
				r.code = v.(int)
			}
		}
	}
	c.JSON(r.httpStatus, gin.H{"code": r.code, "data": r.data, "message": r.message})
	c.Abort()
}
