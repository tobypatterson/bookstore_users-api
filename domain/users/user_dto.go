package users

import (
	"strings"

	"github.com/tobypatterson/bookstore_users-api/utils/errors"
)

// User object represents an user
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Users list of User entities
type Users []User

// Validate will validate the data for the user object
func (user *User) Validate() *errors.RestErr {

	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid Password")
	}

	return nil
}
