package database

import (
	"fmt"
	"log"

	"github.com/ThiagoKufa/ref_api_go/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBInstance é uma instância global do banco de dados
var DBInstance *gorm.DB

// InitDatabase inicializa a conexão com o banco de dados
func ConectDatabase() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar as configurações: %v", err))
	}

	dns := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword)

	fmt.Println(dns)
	DBInstance, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Falha ao conectar com o banco de dados: %v", err))
	}

	fmt.Println("Conexão com o banco de dados estabelecida")

}

func CloseDatabaseConnection() {
	db, err := DBInstance.DB()
	if err != nil {
		log.Fatalf("Erro ao fechar a conexão com o banco de dados: %v", err)
	}
	err = db.Close()
	if err != nil {
		log.Fatalf("Erro ao fechar a conexão com o banco de dados: %v", err)
	}
	log.Println("Conexão com o banco de dados fechada com sucesso")
}

func Migrate(models ...interface{}) {
	// Realiza a migração para cada modelo passado
	for _, model := range models {
		err := DBInstance.AutoMigrate(model)
		if err != nil {
			panic(fmt.Sprintf("Falha na migração do modelo [%v]: %v", model, err))
		}
	}

	fmt.Println("Migração do banco de dados concluída")
}
