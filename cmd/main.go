package main

import (
	"fmt"
	"net/http"

	"communications/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	fmt.Println("Starting server on port " + cfg.Port)
	router.Run(":" + cfg.Port)
}
