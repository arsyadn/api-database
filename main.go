package main

import (
	"auth-api/config"
	"auth-api/models"
	"auth-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeApp() *gin.Engine {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

db := config.ConnectDatabase()

	db.AutoMigrate(&models.User{})

	routes.SetupRoutes(r, db)

	return r
}

func main() {
	app := InitializeApp()
	app.Run(":8080")
}