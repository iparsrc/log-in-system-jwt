package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/parsaakbari1209/simple-log-in-system-jwt/domain"
)

const (
	secret = "secret key" // The secret key must be read from a secure file or environment variable.
)

// NewJwtString creates a JWT based on the user information.
// Returns a JWT string if succed to sign the token.
// Returns a RESTError if falied to sign the token.
func NewJwtString(user *domain.User) (string, *RESTError) {
	// Set JWT claims.
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   strconv.FormatInt(time.Now().Add(time.Minute*5).Unix(), 10),
	}
	// Create and sign the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", &RESTError{
			Message: "can not sign json web token.",
			Status:  500,
			Error:   "internal server error",
		}
	}
	// Return the token(string type) and nil(no error).
	return tokenString, nil
}

// ValidateJwt gets an access-token(JWT) and validates.
// Returns user if JWT is valid and not expired or changed.
// Returns a RESTError if JWT is expired or changed in the client side.
func ValidateJwt(AT string) (*domain.User, *RESTError) {
	// Parse token and validate using signiture.
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
	// Extract all claims in the JWT payload.
	claims, _ := token.Claims.(jwt.MapClaims)
	exp, _ := claims["exp"].(string)
	userId, _ := claims["id"].(string)
	userName, _ := claims["name"].(string)
	userEmail, _ := claims["email"].(string)
	// Change the exp type from string to int64.
	expInt, _ := strconv.ParseInt(exp, 10, 64)
	// Check, if the token in expired or not.
	if time.Now().Unix() > expInt {
		return nil, &RESTError{
			Message: "token is expired.",
			Status:  400,
			Error:   "bad request",
		}
	}
	// Create the user, based on the claims.
	user := domain.User{
		ID:    userId,
		Email: userEmail,
		Name:  userName,
	}
	// Return the user, and nil(no error).
	return &user, nil
}
