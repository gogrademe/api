package handler

import (
	"net/http"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

type Setup struct {
	model.Person
	Email    string
	Password string
}

func SetupApp(c *echo.Context) error {
	db := ToDB(c)
	accts, err := db.GetAccountList()
	if len(accts) > 0 || err != nil {
		return ErrNotAuthorized
	}

	p := &Setup{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.InsertPerson(&p.Person); err != nil {
		return ErrSaving.Log(err)
	}

	u, _ := model.NewAccountFor(p.ID, p.Email)
	if err := u.SetPassword(p.Password); err != nil {
		return ErrPasswordSimple.Log(err)
	}

	u.SetActive()

	if err := db.InsertAccount(u); err != nil {
		return ErrSaving.Log(err)
	}

	return c.JSON(http.StatusCreated, "")
}
