package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	Id        uint64    `json:"id,omitempty`
	Name      string    `json:"nome,omitempty"`
	Email     string    `json:"email,omitempty`
	Password  string    `json:password,omitempty`
	CreatedAt time.Time `json:"CreatedAt,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if erro := user.validate(stage); erro != nil {
		return erro
	}
	if erro := user.format(stage); erro != nil {
		return erro
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("The name can't be empty")
	}
	if user.Email == "" {
		return errors.New("The E-mail can't be empty")
	}
	if stage == "register" && user.Password == "" {
		return errors.New("The password is a required field")
	}
	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New(("Invalid e-mail"))
	}
	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passHash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(passHash)
	}
	return nil
}
