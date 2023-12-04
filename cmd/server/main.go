package main

import (
	"github.com/ThiagoKufa/ref_api_go/internal/entities"
	"github.com/ThiagoKufa/ref_api_go/internal/infra/database"
)

func main() {
	database.ConectDatabase()
	database.Migrate(&entities.Product{})
	database.CloseDatabaseConnection()
}
