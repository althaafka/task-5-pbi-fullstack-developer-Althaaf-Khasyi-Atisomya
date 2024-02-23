package productcontroller

import (
	"net/http"

	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/models"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/database"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var photo models.Photo

	if err:= c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&photo)


	c.JSON(http.StatusCreated, gin.H{"data": photo})
}
