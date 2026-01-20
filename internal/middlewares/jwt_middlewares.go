package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyJwtToken(c *gin.Context) {
	jwtToken := c.GetHeader("Authorization")
	if jwtToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "missing_token",
			"message": "Authorization header is absent or empty",
		})
		return
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET_KEY, nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "invalid_token",
			"message": fmt.Sprintf("JWT validation failed: %v", err),
		})
		return
	}

	c.Next()
}
