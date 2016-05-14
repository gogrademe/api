package model

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidPassword = errors.New("password must be between 6 and 256 characters")

// IsActive checks Disabled and ActivationToken fields.
func (u *Account) IsActive() bool {
	return u.ActivationToken == "" && u.Disabled == false
}

// SetActive clears Disabled and ActivationToken fields.
func (u *Account) SetActive() {
	u.ActivationToken = ""
	u.Disabled = false
}

// GenActivationToken will create a random token to be used to activate the account.
func (u *Account) GenActivationToken() error {
	rb := make([]byte, 32)
	_, err := rand.Read(rb)
	if err != nil {
		return err
	}
	u.ActivationToken = base64.URLEncoding.EncodeToString(rb)
	return nil
}

// NewAccountFor will create a new user with an activation token.
func NewAccountFor(personID int, email string) (*Account, error) {
	account := &Account{
		PersonID: personID,
		Email:    email,
		Disabled: true,
	}
	if err := account.GenActivationToken(); err != nil {
		return nil, err
	}

	return account, nil
}

// SetPassword will validate and hash a password string
func (u *Account) SetPassword(password string) error {
	// Password validation.
	switch {
	case len(password) < 6:
		return ErrInvalidPassword
	case len(password) > 256:
		return ErrInvalidPassword
	}
	// Hash password
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.HashedPassword = string(b)

	return nil
}

// ComparePassword returns an error if password doesn't match password hash
func (u *Account) ComparePassword(password string) error {
	if u.HashedPassword == "" {
		return ErrInvalidPassword
	}
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
}
