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

	var existingUser models.User
	if err := database.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existingUser).Error; err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Username or Email already exists"})
		return
	}

	database.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var user models.User

	if err:= c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User
	if err := database.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&existingUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": existingUser})
}

func Update(c *gin.Context) {
	var user models.User

	if err:= c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := c.Param("userId")
	var existingUser models.User
	if err := database.DB.Where("id = ?", userId).First(&existingUser).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Model(&existingUser).Updates(&user)

	c.JSON(http.StatusOK, gin.H{"data": existingUser})
}