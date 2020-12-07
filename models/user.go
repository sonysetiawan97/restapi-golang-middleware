package models

// User ...
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"email"`
	Password string `json:"password"`
}
