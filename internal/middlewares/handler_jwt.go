package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func VerifyJwtToken(c *gin.Context) {
	jwtToken := c.Request.Header.Get("Authorization")

	if jwtToken == "" {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Missing JWT Token: the Authorization header is absent or empty.",
		})
		return
	}

	token, _ := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JWT_SECRET_KEY, nil
	})

	if token.Valid == false {
		c.Abort()
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": "JWT Token validation error: signature verification failed, invalid signing key",
		})
		return
	}
	c.Next()
}
