package handlers

import (
	"github.com/gin-gonic/gin"
	"gonas/models"
	"gonas/utils"
	"log"
	"net/http"
)

func TestModel(c *gin.Context)  {
	user, err := models.FindUserByID(1)
	if err != nil {
		utils.ErrDetail(err)
	}
	log.Println(user)
	files, err := user.UserFiles()
	log.Println(files)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
