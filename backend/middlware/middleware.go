package middlware

import (
	"github.com/gin-gonic/gin"
	"gonas/handlers"
	"gonas/models"
	"gonas/utils"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Admin-Token")
		if err != nil || !checkAuth(token) {
			c.JSON(http.StatusUnauthorized, gin.H{"please login first": ""})
			handlers.ClientJson(c, http.StatusUnauthorized, "", handlers.CodeInvalidParam, "用户token失效！")
			c.Abort()
		}
		c.Next()
	}
}

func checkAuth(token string) bool {
	decryptToken, err := models.DecryptToken(token)
	if err != nil {
		utils.ErrDetail(err)
		return false
	}
	return decryptToken.ID != 0
}
