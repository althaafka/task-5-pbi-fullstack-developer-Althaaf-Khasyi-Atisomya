package helpers

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// JWT claims structure
type JWTClaims struct {
	UserID uint
	jwt.StandardClaims
}

// GenerateToken generates a JWT token
func GenerateToken(userID uint) (string, error) {
	// Set expiration time of the token
	expTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token
func ValidateToken(tokenString string) (*JWTClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	// Validate the token and return the claims
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
