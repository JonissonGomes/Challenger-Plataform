package repositories

import (
	"log"

	"github.com/JonissonGomes/Challenger-Plataform/domain"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}

// Struct para armazenar os dados no banco
type UserRepositoryDataBase struct {
	Db *gorm.DB
}

func (repo UserRepositoryDataBase) Insert(user *domain.User) (*domain.User, error) {
	err := user.Prepare()

	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
	}

	// Persistindo dados no banco
	err = repo.Db.Create(user).Error

	if err != nil {
		log.Fatalf("Error to persist user on database: %v", err)
		return user, err
	}

	// Retornando resultado da operação
	return user, nil
}
