package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/app/config"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/app/models"
)

func Start() {
	router := gin.Default()

	db, err := config.SetupDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(
		&models.Class{},
		&models.ClassCategory{},
		&models.User{},
		&models.UserType{},
		&models.UserClass{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
