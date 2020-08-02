package utils

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/parsaakbari1209/simple-log-in-system-jwt/domain"
)

const (
	secret = "secret key" // The secret key must be in a secure file or environment variable.
)

func NewJwtString(user *domain.User) (string, *RESTError) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().UTC().Add(time.Minute * 2).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", &RESTError{
			Message: "can not create json web token.",
			Status:  500,
			Error:   "internal server error",
		}
	}
	return tokenString, nil
}
