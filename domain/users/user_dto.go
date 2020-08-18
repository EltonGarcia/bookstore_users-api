package users

import (
	"github.com/EltonGarcia/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64 `json:"id"`
	FistName    string `json:"fist_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == ""{
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}