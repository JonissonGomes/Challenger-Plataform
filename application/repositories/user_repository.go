package repositories

import "github.com/JonissonGomes/Challenger-Plataform/domain"

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
}
