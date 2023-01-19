package main

import (
	"fmt"
	"log"
	"net/http"

	"google-auth/internal/configs"
	"google-auth/internal/loger"
	"google-auth/internal/services"

	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper across the application
	configs.InitializeViper()

	// Initialize Logger across the application
	loger.InitializeZapCustomLogger()

	// Initialize Oauth2 Services
	services.InitializeOAuthGoogle()

	// Routes for the application
	http.HandleFunc("/", services.HandleMain)
	http.HandleFunc("/login-gl", services.HandleGoogleLogin)
	http.HandleFunc("/callback-gl", services.CallBackFromGoogle)

	// loger.Log.Info("Started running on http://localhost:" + viper.GetString("port"))
	fmt.Println("welcome")
	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))
}
