package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func TestModel(c *gin.Context)  {
	var s []string
	s = append(s, "aaaa")
	log.Println(s)
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
