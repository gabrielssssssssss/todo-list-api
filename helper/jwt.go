package helper

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Id    string
	Email string
	jwt.RegisteredClaims
}

func GenerateJwtToken(id string, email string, key []byte) (string, error) {
	claims := JWTClaim{
		Id:    id,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 300)),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Unable to create JWT with the given parameters.")
	}

	return token, nil
}
