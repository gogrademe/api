package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// NewSession will create a jwt token for the user after we verified their password.s
func NewSession(u Account) (Session, error) {
	var s Session
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	tokenExpires := time.Now().UTC().Add(time.Hour * 72)
	token.Claims["userID"] = u.ID
	token.Claims["personID"] = u.PersonID
	token.Claims["email"] = u.Email
	token.Claims["exp"] = tokenExpires.Unix()

	// TODO: Move this to a config file.
	tokenString, err := token.SignedString([]byte("someRandomSigningKey"))
	if err != nil {
		return s, err
	}

	s = Session{
		Token:     tokenString,
		AccountID: u.ID,
		ExpiresAt: tokenExpires,
	}

	return s, err
}
