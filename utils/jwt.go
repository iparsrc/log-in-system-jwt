package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/parsaakbari1209/simple-log-in-system-jwt/domain"
)

const (
	secret = "secret key" // The secret key must be in a secure file or environment variable.
)

func NewJwtString(user *domain.User) (string, *RESTError) {
	claims := jwt.MapClaims{
		"id":    strconv.FormatInt(user.ID, 10),
		"name":  user.Name,
		"email": user.Email,
		"exp":   strconv.FormatInt(time.Now().Add(time.Minute*2).Unix(), 10),
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
	exp, _ := claims["exp"].(string)
	userId, _ := claims["id"].(string)
	userEmail, _ := claims["email"].(string)
	userName, _ := claims["name"].(string)

	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	expInt, _ := strconv.ParseInt(exp, 10, 64)

	if time.Now().Unix() > expInt {
		return nil, &RESTError{
			Message: "token is expired.",
			Status:  400,
			Error:   "bad request",
		}
	}

	user := domain.User{
		ID:    userIdInt,
		Email: userEmail,
		Name:  userName,
	}

	return &user, nil
}
