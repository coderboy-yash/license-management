package main

import (
	"log"

	"license-service/api/handler"
	"license-service/config"
	"license-service/repository"
	"license-service/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	config.ConnectDB()

	repo := repository.NewLicenseRepository()
	svc := service.NewLicenseService(repo)
	h := handler.NewLicenseHandler(svc)

	r := gin.Default()

	r.POST("/licenses", h.CreateLicense)
	r.GET("/licenses/:license_id/features", h.GetLicenseFeatures)

	r.Run(":8080")
}
