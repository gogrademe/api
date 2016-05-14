package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

var (
	ErrPasswordSimple         = NewAPIError(http.StatusBadRequest, "password too simple")
	ErrInvalidActivationToken = NewAPIError(http.StatusBadRequest, "activation token does not match any accounts")
	ErrSaving                 = NewAPIError(http.StatusInternalServerError, "error saving record")
	ErrNotAuthorized          = NewAPIError(http.StatusUnauthorized, "access denied")
)

// GetAllAccounts http endpoint to return all accounts.
func GetAllAccounts(c echo.Context) error {
	db := ToDB(c)

	r, err := db.GetAccountList()
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, r)

}

func CreateAccount(c echo.Context) error {
	a := &model.Account{}
	if err := c.Bind(a); err != nil {
		return ErrBind.Log(err)
	}

	na, err := model.NewAccountFor(a.PersonID, a.Email)
	if err != nil {
		return ErrServerError.Log(err)
	}

	db := ToDB(c)
	if err := db.InsertAccount(na); err != nil {
		return ErrSaving.Log(err)
	}

	return c.JSON(http.StatusCreated, na)
}

func DeleteAccount(c echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DeleteAccount(id); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusOK, nil)
}

// ActivateAccount will activate a user account from a token or an admin.
func ActivateAccount(c echo.Context) error {
	db := ToDB(c)
	token, password := c.Param("token"), c.FormValue("password")
	usr, err := db.GetAccountByToken(token)
	if err != nil {
		return ErrInvalidActivationToken.Log(err)
	}

	usr.SetActive()
	if err := usr.SetPassword(password); err != nil {
		return ErrPasswordSimple.Log(err)
	}

	if err := db.UpdateAccount(usr); err != nil {
		return ErrSaving.Log(err)
	}

	return c.JSON(http.StatusOK, usr)
}
