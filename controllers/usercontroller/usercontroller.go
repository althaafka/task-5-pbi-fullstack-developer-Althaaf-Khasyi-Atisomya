package usercontroller

import (
	"net/http"

	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/models"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/database"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user models.User

	if err:= c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//check if username or email already exists
	var existingUser models.User
	if err := database.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Username or Email already exists"})
		return
	}

	database.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}