package models

import (
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	PaymentID   uuid.UUID `json:"payment_id"`
	PaymentDate time.Time `json:"payment_date"`
	CustomerID  uuid.UUID `json:"customer_id"`
	Customer    Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
}
