package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `gorm:"type:varchar(100);not null" json:"name"`
	Description string     `gorm:"type:text;not null" json:"description"`
	Price       float64    `gorm:"not null" json:"price"`
	Quantity    int        `gorm:"not null" json:"quantity"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`
	CreatedBy   string     `gorm:"type:varchar(100);not null" json:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}
