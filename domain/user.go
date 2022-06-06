package domain

type User struct {
	BaseStructModel
	Name     string
	Email    string
	Password string
	TokenId  string
}
