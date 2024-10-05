package models

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID          uuid.UUID `gorm:"type:char(36);primary_key;" json:"id"`
	FirstName   string    `gorm:"type:varchar(255)" json:"first_name"`
	LastName    string    `gorm:"type:varchar(255)" json:"last_name"`
	UserID      uuid.UUID `gorm:"type:char(36);index;" json:"user_id"`
	UserAccount User      `gorm:"foreignKey:UserID" json:"user_account"`
}
