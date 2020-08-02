package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
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
		"exp":   time.Now().UTC().Add(2 * time.Hour).Unix(),
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

func ValidateJwt(AT string) (*domain.User, *RESTError) {
	token, err := jwt.Parse(AT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, &RESTError{
			Message: "can not parse jwt.",
			Status:  400,
			Error:   "bad request",
		}
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	exp, _ := claims["exp"].(int64)
	userId, _ := claims["id"].(int64)
	userEmail, _ := claims["id"].(string)
	if time.Now().UTC().After(time.Unix(exp, 0)) {
		return nil, &RESTError{
			Message: "token is expired.",
			Status:  400,
			Error:   "bad request",
		}
	}
	user := domain.User{
		ID:    userId,
		Email: userEmail,
	}
	return &user, nil
}
