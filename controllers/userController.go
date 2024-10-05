package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"test-mnc/initializers"
	"test-mnc/models"
	"time"
)

func SignUp(c *gin.Context) {
	// Request body
	var body struct {
		Email     string
		Password  string
		FirstName string
		LastName  string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Encrypt password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Start a new transaction
	err = initializers.DB.Transaction(func(tx *gorm.DB) error {
		// Create User
		user := models.User{
			ID:       uuid.New(),
			Email:    body.Email,
			Password: string(hash),
		}
		result := tx.Create(&user)
		if result.Error != nil {
			return result.Error
		}

		// Create Customer
		customer := models.Customer{
			ID:        uuid.New(),
			FirstName: body.FirstName,
			LastName:  body.LastName,
			UserID:    user.ID,
		}
		result = tx.Create(&customer)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	// Handle transaction errors
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create User and Customer",
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"first_name": body.FirstName,
		"last_name":  body.LastName,
		"email":      body.Email,
	})
}

func Login(c *gin.Context) {
	// Request body
	var body struct {
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Checking User
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == uuid.Nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User not found",
		})
		return
	}

	// Checking password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect password",
		})
		return
	}

	// Start a new transaction
	err = initializers.DB.Transaction(func(tx *gorm.DB) error {
		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			return err
		}

		// Update User's token
		user.Token = tokenString
		result := tx.Save(&user)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	// Handle transaction errors
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token and update user",
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"token": user.Token,
	})
}
