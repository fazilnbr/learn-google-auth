package main

import (
	"google-auth/internal/configs"
	"google-auth/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Oauth2 Services
	services.InitializeOAuthGoogle()

	engine := gin.New()
	engine.LoadHTMLGlob("*.html")
	engine.Use(gin.Logger())
	// Use logger from Gin

	engine.GET("/", services.HandleMain)
	engine.GET("/login-gl", services.GoogleLogin)
	engine.GET("/callback-gl", services.CallBackFromGoogle)

	engine.Run()

}
