package models

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	IsAmbassador bool   `json:"isambassador"`
}
