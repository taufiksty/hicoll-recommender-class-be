package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/config"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/handlers"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/middlewares"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/models"
	"gorm.io/gorm"
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

	setupRoutes(router, db)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(router *gin.Engine, db *gorm.DB) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.POST("/api/user/register", func(ctx *gin.Context) {
		handlers.Register(ctx, db)
	})
	router.POST("/api/user/login", func(ctx *gin.Context) {
		handlers.Login(ctx, db)
	})

	authRouter := router.Use(middlewares.AuthMiddleware())
	authRouter.PUT("/api/user/:id", func(ctx *gin.Context) {
		handlers.UpdateUser(ctx, db)
	})
	authRouter.GET("/api/class", func(ctx *gin.Context) {
		handlers.GetClassByCategory(ctx, db)
	})
	authRouter.GET("/api/class/recommendation", func(ctx *gin.Context) {
		handlers.GetRecommendationHandlers(ctx, db)
	})
}
