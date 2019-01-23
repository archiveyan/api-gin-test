package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weet/model"
)

func GetAnswers(c *gin.Context) {

	answers, err := model.GetAnswers()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, answers)
}
