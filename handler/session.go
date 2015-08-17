package handler

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

const bearer = "Bearer"

var ErrInvalidCredentials = NewAPIError(http.StatusUnauthorized, "invalid email and/or password")

// LoginForm only used for retrieving login credentials.
type LoginForm struct {
	Email    string
	Password string
}

// CreateSession retrieves a user account, checks if active and compares hashed
// password with provided password.
func CreateSession(key, method string) echo.HandlerFunc {
	return func(c *echo.Context) error {
		p := &LoginForm{}
		if err := c.Bind(p); err != nil {
			return ErrBind.Log(err)
		}

		db := ToDB(c)
		account, err := db.GetAccountEmail(p.Email)
		if err != nil {
			return ErrInvalidCredentials.Log(err)
		}

		if !account.IsActive() {
			return c.JSON(http.StatusUnauthorized, "account inactive")
		}
		if err := account.ComparePassword(p.Password); err != nil {
			return ErrNotAuthorized.Log(err)
		}

		session, err := model.NewSession(key, method, *account)
		if err != nil {
			return ErrNotAuthorized.Log(err)
		}

		if err := db.InsertSession(&session); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, session)

	}
}

// JWTAuth is a JSON Web Token middleware
func JWTAuth(key, method string) echo.HandlerFunc {
	return func(c *echo.Context) error {

		// Skip WebSocket
		if (c.Request().Header.Get(echo.Upgrade)) == echo.WebSocket {
			return nil
		}

		auth := c.Request().Header.Get("Authorization")
		l := len(bearer)

		if len(auth) > l+1 && auth[:l] == bearer {
			t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {

				// Always check the signing method
				expected := jwt.GetSigningMethod(method)
				if token.Method != expected {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// Return the key for validation
				return []byte(key), nil
			})
			if err == nil && t.Valid {
				// Store token claims in echo.Context
				c.Set("claims", t.Claims)
				return nil
			}
		}
		return ErrNotAuthorized.Log(nil)
	}
}
