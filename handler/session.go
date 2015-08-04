package handler

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gogrademe/api/model"
	"github.com/gogrademe/api/store"
	"github.com/labstack/echo"
)

const bearer = "Bearer"

type LoginForm struct {
	Email    string
	Password string
}

func CreateSession(c *echo.Context) error {
	p := &LoginForm{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db := ToDB(c)
	account, err := store.GetAccountEmail(db, p.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	if !account.IsActive() {
		return c.JSON(http.StatusUnauthorized, "account inactive")
	}
	if err := account.ComparePassword(p.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	session, err := model.NewSession(*account)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	if err := store.InsertSession(db, &session); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, &M{"session": session})

}

// A JSON Web Token middleware
func JWTAuth(key string) echo.HandlerFunc {
	return func(c *echo.Context) error {

		// Skip WebSocket
		if (c.Request().Header.Get(echo.Upgrade)) == echo.WebSocket {
			return nil
		}

		auth := c.Request().Header.Get("Authorization")
		l := len(bearer)
		he := echo.NewHTTPError(http.StatusUnauthorized)

		if len(auth) > l+1 && auth[:l] == bearer {
			t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {

				// Always check the signing method
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
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
		return he
	}
}
