package handlers

import (
	"github.com/gin-gonic/gin"
	"gonas/models"
	"gonas/utils"
	"log"
	"net/http"
)

func TestModel(c *gin.Context) {
	user, _ := models.FindUserByID(1)
	files, _ := models.FilesByUser(&user)
	log.Println("files: ", files)
	utils.ClientJson(c, http.StatusOK, files, utils.CodeSuccess, "ok")
}
