package domain

import (
	"log"

	uuid "github.com/stori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseStructModel
	Name     string
	Email    string
	Password string
	TokenId  string
}

func (user *User) Prepare() error {

	// Encryptando senha
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	// Tratando erros
	if err != nil {
		log.Fatalf("Error in duration password generation: %v", err)
	}

	// Atribuindo senha encriptada
	user.Password = string(password)

	// Gerando Token automático
	user.TokenId = uuid.NewV4().String()
}
