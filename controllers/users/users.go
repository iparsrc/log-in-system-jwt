package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parsaakbari1209/simple-log-in-system-jwt/domain"
	"github.com/parsaakbari1209/simple-log-in-system-jwt/services"
	"github.com/parsaakbari1209/simple-log-in-system-jwt/utils"
)

func CreateUser(c *gin.Context) {
	var newUser domain.User
	// Validate the givin JSON formated data.
	if err := c.ShouldBindJSON(&newUser); err != nil {
		restErr := utils.BadRequest("Invalid JSON Sent to create a new user.")
		c.JSON(restErr.Status, restErr)
		return
	}
	// Create the new user.
	AT, restErr := services.Create(&newUser)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}
	// Responde with the created uesr profile.
	c.SetCookie("access_token", AT, 1000, "/", "localhost:8080", false, false)
	c.JSON(http.StatusOK, gin.H{
		"id":    newUser.ID,
		"email": newUser.Email,
		"name":  newUser.Name,
	})
}
