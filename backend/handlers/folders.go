package handlers

import (
	"github.com/gin-gonic/gin"
	"gonas/models"
	"gonas/utils"
	"log"
	"net/http"
)

func CreateFolder(c *gin.Context)  {
	var folder models.Folder
	err := c.BindJSON(&folder)
	if err != nil {
		utils.ErrDetail(err)
		utils.ClientJson(c, http.StatusBadRequest, "", utils.CodeSuccess, "ok")
		return
	}
	log.Println("name: ", folder.Display)
	utils.ClientJson(c, http.StatusOK, "", utils.CodeSuccess, "ok")
}
