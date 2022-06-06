package domain

import "time"

type BaseStructModel struct {
	ID        string
	CreatedAt time.Timer
	UpdateAt  time.Timer
	DeletedAt time.Timer
}
