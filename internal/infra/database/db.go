package database

import (
	"fmt"
	"log"

	"github.com/ThiagoKufa/ref_api_go/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConectDatabase inicializa e retorna a conexão com o banco de dados
func ConectDatabase() *gorm.DB {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalf("Erro ao carregar as configurações: %v", err)
	}

	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar com o banco de dados: %v", err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida")
	return db
}
