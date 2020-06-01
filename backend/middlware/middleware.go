package middlware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gonas/models"
	"gonas/utils"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := models.CheckAuth(c)
		if err != nil{
			fmt.Println(err)
			utils.ErrDetail(err)
			utils.ClientJson(c, http.StatusUnauthorized, "", utils.CodeInvalidParam, "用户token失效！")
			c.Abort()
		}
		c.Next()
	}
}


