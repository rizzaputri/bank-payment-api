package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	Email     string    `gorm:"type:varchar(255);unique" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	Token     string    `gorm:"type:varchar(255)" json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
