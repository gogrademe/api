package model

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/jinzhu/gorm"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("password must be between 6 and 256 characters")
)

type User struct {
	gorm.Model

	PersonID        uint `sql:"index"`
	Role            string
	Email           string `sql:"type:varchar(100);unique_index"`
	HashedPassword  string `json:"-"`
	ActivationToken string `json:"-"` // base64 url encoded random hash.
	Disabled        bool
}

// GenActivationToken will create a random token to be used to activate the account.
func (u *User) GenActivationToken() error {
	rb := make([]byte, 32)
	_, err := rand.Read(rb)
	if err != nil {
		return err
	}
	u.ActivationToken = base64.URLEncoding.EncodeToString(rb)
	return nil
}

func NewUserFor(email string, personID uint) (*User, error) {
	user := &User{
		Email:    email,
		PersonID: personID,
		Disabled: true,
	}
	if err := user.GenActivationToken(); err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserForWithPassword(email, password string, personID uint) (*User, error) {
	user := &User{
		Email:    email,
		PersonID: personID,
		Disabled: false,
	}

	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) SetPassword(password string) error {
	// Password validation.
	switch {
	case len(password) < 6:
		return ErrInvalidPassword
	case len(password) > 265:
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
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
}
