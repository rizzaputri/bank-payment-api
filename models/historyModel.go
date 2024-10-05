package models

import (
	"github.com/google/uuid"
	"time"
)

type History struct {
	HistoryID  uuid.UUID `gorm:"primary_key" json:"history_id"`
	Date       time.Time `json:"date"`
	Activity   string    `gorm:"type:text" json:"activity"`
	CustomerID uuid.UUID `json:"customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"customer"`
}
