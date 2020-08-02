package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {
	MapURLs()
	router.Run(":8080")
}
