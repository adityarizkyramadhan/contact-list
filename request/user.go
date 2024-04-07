package request

import (
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u UserRegister) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Username, validation.Length(8, 50)),
		validation.Field(&u.Username, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$"))),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Password, validation.Length(8, 50)),
		validation.Field(&u.Password, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$"))),
	)
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u UserLogin) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Username, validation.Length(8, 50)),
		validation.Field(&u.Username, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$"))),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Password, validation.Length(8, 50)),
		validation.Field(&u.Password, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$"))),
	)
}

type UserUpdate struct {
	ID       string `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u UserUpdate) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Username, validation.Length(8, 50)),
		validation.Field(&u.Username, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$"))),
		validation.Field(&u.Password, validation.Length(8, 50)),
		validation.Field(&u.Password, validation.Match(regexp.MustCompile("^[a-zA-Z0-9]*$"))),
	)
}
