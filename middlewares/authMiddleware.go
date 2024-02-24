package middlewares

import (
    "net/http"
    "strings"
    "github.com/gin-gonic/gin"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/helpers"
)

// AuthMiddleware adalah middleware untuk menangani autentikasi
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Get the client token from the request header
        bearerToken := c.GetHeader("Authorization")
        if bearerToken == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
            c.Abort()
            return
        }

        tokenString := strings.TrimPrefix(bearerToken, "Bearer ")

        // Validate the token
        claims, err := helpers.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Add the userID to the context
        c.Set("userID", claims.UserID)

        c.Next()
    }
}
