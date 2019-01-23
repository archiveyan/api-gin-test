package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weet/model"
)

func CreateSample(c *gin.Context) {
	var sample model.Sample
	c.BindJSON(&sample)

	model.CreateSample(model.Sample{
		Name:    sample.Name,
		Comment: sample.Comment,
	})

	c.JSON(http.StatusOK, sample)
}

func GetSampleById(c *gin.Context) {
	id, err := GetUint(c, "sample_id")
	content, err := model.GetSampleById(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, content)
}
