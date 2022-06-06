package domain

import (
	"log"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseStructModel
	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255);unique_index"`
	Password string `gorm:"type:varchar(255)"`
	TokenId  string `gorm:"type:varchar(255);unique_index"`
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
