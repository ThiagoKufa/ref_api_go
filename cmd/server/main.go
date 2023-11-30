package main

import (
	"time"

	"github.com/ThiagoKufa/ref_api_go/internal/infra/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // Inclui campos padrões do GORM como ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	FirstName  string
	LastName   string
	BirthDate  time.Time
	IsActive   bool `gorm:"default:true"`
}

func main() {
	database.ConectDatabase()
	database.Migrate(&User{})
	database.CloseDatabaseConnection()

}
