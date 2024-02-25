package usercontroller

import (
	"net/http"

	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/models"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/database"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/helpers"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User

	if err:= c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)

	database.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var user models.User

	if err:= c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid login credentials"})
		return
	}

    token, err := helpers.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"id":user.ID, "token": token})

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

func Delete(c *gin.Context) {
	userId := c.Param("userId")
	var user models.User
	if err := database.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}