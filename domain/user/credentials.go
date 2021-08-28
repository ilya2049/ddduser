package user

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var (
	ErrNameRequired     = errors.New("user name is required")
	ErrEmailRequired    = errors.New("user email is required")
	ErrInvalidEmail     = errors.New("user email is invalid")
	ErrPasswordRequired = errors.New("user password is required")
	ErrPasswordTooShort = errors.New("user password is too short")
)

func NewCredentials(
	name string,
	email string,
) (Credentials, error) {
	if name == "" {
		return Credentials{}, ErrNameRequired
	}

	if email == "" {
		return Credentials{}, ErrEmailRequired
	}

	if !strings.Contains(email, "@") {
		return Credentials{}, ErrInvalidEmail
	}

	return Credentials{
		name:  name,
		email: email,
	}, nil
}

type Credentials struct {
	name     string
	password string
	email    string
}

func (c *Credentials) AddPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	const passwordMinLength = 6

	if utf8.RuneCountInString(password) < passwordMinLength {
		return ErrPasswordTooShort
	}

	c.password = password

	return nil
}

func (c Credentials) Name() string {
	return c.name
}

func (c Credentials) Password() string {
	return c.password
}

func (c Credentials) Email() string {
	return c.email
}
