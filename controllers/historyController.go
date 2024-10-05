package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test-mnc/initializers"
	"test-mnc/models"
)

func GetHistory(c *gin.Context) {
	var histories []models.History

	// Query all payments from the database
	result := initializers.DB.Preload("Customer").Find(&histories)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to retrieve histories",
		})
		return
	}

	// Response with the list of payments
	c.JSON(http.StatusOK, gin.H{
		"histories": histories,
	})
}
