package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// NewSession will create a jwt token for the user after we verified their password.s
func NewSession(key, method string, u Account) (Session, error) {
	var s Session
	token := jwt.New(jwt.GetSigningMethod(method))

	tokenExpires := time.Now().UTC().Add(time.Hour * 72)
	token.Claims["account_id"] = u.AccountID
	token.Claims["person_id"] = u.PersonID
	token.Claims["email"] = u.Email
	token.Claims["exp"] = tokenExpires.Unix()

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return s, err
	}

	s = Session{
		Token:     tokenString,
		AccountID: u.AccountID,
		ExpiresAt: tokenExpires,
	}

	return s, err
}
