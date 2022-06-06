package main

import (
	"fmt"
	"log"

	"github.com/JonissonGomes/Challenger-Plataform/application/repositories"
	"github.com/JonissonGomes/Challenger-Plataform/domain"
	"github.com/JonissonGomes/Challenger-Plataform/framework/utils"
)

func main() {

	// Inicializando conexão
	db := utils.ConnectDatabase()

	// Criando usuário
	user := domain.User{
		Name:     "Jonisson Gomes",
		Email:    "jonis@gmail.com",
		Password: "123456789",
	}

	// Estabelecendo relação como banco
	userRepo := repositories.UserRepositoryDataBase{Db: db}

	result, err := userRepo.Insert(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
