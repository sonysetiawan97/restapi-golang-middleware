package models

// User ...
type User struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// TableUser ...
func (b *User) TableUser() string {
	return "user"
}
