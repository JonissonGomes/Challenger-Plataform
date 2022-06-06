package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
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
		return err
	}

	// Atribuindo senha encriptada
	user.Password = string(password)

	// Gerando Token automático
	user.TokenId = uuid.NewV4().String()

	err = user.validateField()

	if err != nil {
		log.Fatalf("Error during the user validate: %v", err)
		return err
	}

	return nil
}

func (user User) validateField() error {
	return nil
}
