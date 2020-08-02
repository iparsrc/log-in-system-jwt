package app

import "github.com/parsaakbari1209/simple-log-in-system-jwt/controllers/ping"

func MapURLs() {
	router.GET("/ping", ping.Ping)
}
