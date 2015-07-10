package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Session struct {
	gorm.Model
	Token     string    `json:"token"`
	UserID    uint      `json:"userId"`
	ExpiresAt time.Time `json:"expiresAt"`
}

// NewSession will create a jwt token for the user after we verified their password.
func NewSession(u User) (Session, error) {
	var s Session
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	tokenExpires := time.Now().UTC().Add(time.Hour * 72)
	token.Claims["userId"] = u.ID
	token.Claims["personId"] = u.PersonID
	token.Claims["email"] = u.Email
	token.Claims["exp"] = tokenExpires.Unix()

	// TODO: Move this to a config file.
	tokenString, err := token.SignedString([]byte("someRandomSigningKey"))
	if err != nil {
		return s, err
	}

	s = Session{
		Token:     tokenString,
		UserID:    u.ID,
		ExpiresAt: tokenExpires,
	}

	return s, err
}
