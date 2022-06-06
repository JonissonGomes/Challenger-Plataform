package domain

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BaseStructModel struct {
	ID        string     `json:"id" gorm:"type:uuid; primary_key"`
	CreatedAt time.Timer `json:"created_at" gorm:"type:datetime"`
	UpdateAt  time.Timer `json:"updated_at" gorm:"type:datetime"`
	DeletedAt time.Timer `json:"deleted_at" gorm:"type:datetime" sql:"index"`
}

func CreateNewUser() *User {
	return &User{}
}

func (base *BaseStructModel) BeforCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil {
		log.Fatalf("Erro during create object")
	}

	err = scope.SetColumn("CreatedAt", time.Now())

	if err != nil {
		log.Fatalf("Erro during create object")
	}

	return nil
}
