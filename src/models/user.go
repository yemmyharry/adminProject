package models

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
	//ConfirmPassword []byte `json:"confirm_password"`
	IsAmbassador bool `json:"is_ambassador"`
}
