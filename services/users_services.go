package services

import (
	"github.com/parsaakbari1209/simple-log-in-system-jwt/domain"
	"github.com/parsaakbari1209/simple-log-in-system-jwt/utils"
)

func Create(user *domain.User) (string, *utils.RESTError) {
	if (domain.Storage[user.ID] != domain.User{}) {
		return "", &utils.RESTError{
			Message: "user exists",
			Status:  400,
			Error:   "bad request",
		}
	}
	domain.Storage[user.ID] = *user
	AT, restErr := utils.NewJwtString(user)
	if restErr != nil {
		return "", restErr
	}
	return AT, nil
}
