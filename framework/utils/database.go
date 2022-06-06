package utils

import (
	"log"
	"os"

	"github.com/JonissonGomes/Challenger-Plataform/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDatabase() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Lê o arquivo no .env
	readConnectionDatabase := os.Getenv("dsn")

	// Conecta com o banco
	connectDatabase, err := gorm.Open("postgres", readConnectionDatabase)

	// Valida se a conexão foi estabelecida
	if err != nil {
		log.Fatalf("Error connecting to database : %v", err)
		panic(err)
	}

	// Cria as tabelas automaticamente, baseado na estrutura passada como parametro
	connectDatabase.AutoMigrate(&domain.User{})

	return connectDatabase
}
