package middlewares

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var (
	JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
)

type JWTClaim struct {
	Id    string
	Email string
	jwt.RegisteredClaims
}

func GetEmail(claims *JWTClaim) string {
	return claims.Email
}

func GetClaims(claims *JWTClaim) *JWTClaim {
	return claims
}
