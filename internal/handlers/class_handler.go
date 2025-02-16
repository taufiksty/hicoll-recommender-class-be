package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taufiksty/hicoll-recommender-class-be/client"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/models"
	"gorm.io/gorm"
)

func GetRecommendationHandlers(c *gin.Context, db *gorm.DB) {
	userIDStr := c.Query("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	// Call the gRPC client
	recommendations, err := client.GetRecommendations(int32(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"recommendations": recommendations})
}

func GetClassByCategory(ctx *gin.Context, db *gorm.DB) {
	category := ctx.Query("category_id")

	var classes []models.Class

	if err := db.Where("class_category_id = ?", category).Find(&classes).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch classes"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"classes": classes})

}
