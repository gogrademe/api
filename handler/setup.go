package handler

import (
	"fmt"
	"net/http"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

type Setup struct {
	model.Person
	Email    string
	Password string
}

// CanSetup returns StatusOK if able to proceed with setup
func CanSetup(c echo.Context) error {
	db := ToDB(c)
	accts, err := db.GetAccountList()
	if len(accts) > 0 || err != nil {
		return ErrForbidden.Log(err, fmt.Sprintf("total accounts: %s", len(accts)))
	}

	return c.NoContent(200)
}

// SetupApp allows the creation of a admin account if no accounts exist.
func SetupApp(c echo.Context) error {
	db := ToDB(c)
	accts, err := db.GetAccountList()
	if len(accts) > 0 || err != nil {
		return ErrForbidden.Log(err, fmt.Sprintf("total accounts: %s", len(accts)))
	}

	p := &Setup{}
	if err := c.Bind(p); err != nil {
		return ErrBind.Log(err)
	}

	p.IsAdmin = true
	if err := db.InsertPerson(&p.Person); err != nil {
		return ErrSaving.Log(err)
	}

	u, _ := model.NewAccountFor(p.PersonID, p.Email)
	if err := u.SetPassword(p.Password); err != nil {
		return ErrPasswordSimple.Log(err)
	}

	u.SetActive()

	if err := db.InsertAccount(u); err != nil {
		return ErrSaving.Log(err)
	}

	return c.NoContent(http.StatusCreated)
}
