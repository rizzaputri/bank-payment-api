package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"test-mnc/initializers"
	"test-mnc/models"
	"time"
)

func CreatePayment(c *gin.Context) {
	// Request body
	var body struct {
		CustomerID string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Start a new payment
	var payment models.Payment
	var customer models.Customer
	err := initializers.DB.Transaction(func(tx *gorm.DB) error {
		// Checking Customer
		initializers.DB.First(&customer, "id = ?", body.CustomerID)
		if customer.ID == uuid.Nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Customer not found",
			})
			return errors.New("customer not found")
		}

		// Create Payment
		payment = models.Payment{
			PaymentID:   uuid.New(),
			PaymentDate: time.Now(),
			CustomerID:  customer.ID,
			Customer:    customer,
		}
		result := tx.Create(&payment)
		if result.Error != nil {
			return result.Error
		}

		// Log History
		history := models.History{
			HistoryID:  uuid.New(),
			Date:       payment.PaymentDate,
			Activity:   "Successful payment",
			CustomerID: customer.ID,
			Customer:   customer,
		}
		result = tx.Create(&history)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	// Handle payment errors
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create Payment",
		})
		return
	}

	// Response
	c.JSON(http.StatusOK, gin.H{
		"payment_id":   payment.PaymentID,
		"payment_date": payment.PaymentDate,
		"customer_id":  payment.CustomerID,
		"customer": gin.H{
			"first_name": customer.FirstName,
			"last_name":  customer.LastName,
		}})
}

func GetPayment(c *gin.Context) {
	var payments []models.Payment

	// Query all payments from the database
	result := initializers.DB.Preload("Customer").Find(&payments)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to retrieve payments",
		})
		return
	}

	// Response with the list of payments
	c.JSON(http.StatusOK, gin.H{
		"payments": payments,
	})
}
