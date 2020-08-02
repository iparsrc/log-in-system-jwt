package app

import (
	"github.com/parsaakbari1209/simple-log-in-system-jwt/controllers/ping"
	"github.com/parsaakbari1209/simple-log-in-system-jwt/controllers/users"
)

func MapURLs() {
	router.GET("/ping", ping.Ping)
	router.POST("/users/create", users.CreateUser)
}
