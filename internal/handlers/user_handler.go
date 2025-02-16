package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/middlewares"
	"github.com/taufiksty/hicoll-recommender-class-be/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context, db *gorm.DB) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	user.UserTypeID = 3

	if err := db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(ctx *gin.Context, db *gorm.DB) {
	var user models.User
	var input models.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := middlewares.GenerateToken(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func UpdateUser(ctx *gin.Context, db *gorm.DB) {
	userID := ctx.Param("id")

	var updateData struct {
		Interests *string `json:"interests,omitempty"`
	}

	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Interests = updateData.Interests

	if err := db.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}
