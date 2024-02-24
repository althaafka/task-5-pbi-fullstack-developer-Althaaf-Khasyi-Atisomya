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

func Show(c *gin.Context) {
	var photo []models.Photo

	database.DB.Find(&photo)

	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func Update(c *gin.Context) {
	var photo models.Photo

	if err:= c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	photoId := c.Param("photoId")
	var existingPhoto models.Photo
	if err := database.DB.Where("id = ?", photoId).First(&existingPhoto).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	database.DB.Model(&existingPhoto).Updates(&photo)

	c.JSON(http.StatusOK, gin.H{"data": existingPhoto})
}

func Delete(c *gin.Context) {
	photoId := c.Param("photoId")
	var photo models.Photo
	if err := database.DB.Where("id = ?", photoId).First(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	database.DB.Delete(&photo)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
