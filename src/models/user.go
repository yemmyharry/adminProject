package models

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email" gorm:"unique`
	Password     []byte `json:"password"`
	IsAmbassador bool   `json:"-"`
}
