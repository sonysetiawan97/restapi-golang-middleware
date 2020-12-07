package validation

import (
	"regexp"

	"example.com/v1/models"
)

// emailRegex ...
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

// RegisterUser ...
func RegisterUser(user *models.User) (response *string) {
	username := &user.Username
	password := &user.Password

	checkUsername := *username
	lenUsername := len(*username)
	lenPassword := len(*password)

	if !isEmailValid(checkUsername) {
		message := "Email not valid"
		return &message
	}

	if lenUsername == 0 {
		message := "username cannot be empty"
		return &message
	}

	if lenPassword == 0 {
		message := "password cannot be empty"
		return &message
	}

	if len(*username) < 5 {
		message := "username must more or equal 5 words"
		return &message
	}

	if len(*password) < 8 {
		message := "password must more or equal 8 words"
		return &message
	}

	return nil
}
