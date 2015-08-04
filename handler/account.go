package handler

import (
	"log"
	"net/http"

	"github.com/gogrademe/api/model"
	"github.com/gogrademe/api/store"
	"github.com/labstack/echo"
)

// GetAllAccounts http endpoint to return all accounts.
func GetAllAccounts(c *echo.Context) error {
	db := ToDB(c)

	r, err := store.GetAccountList(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, &M{"account": r})

}

func CreateAccount(c *echo.Context) error {
	a := &model.Account{}
	if err := c.Bind(a); err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	na, err := model.NewAccountFor(a.PersonID, a.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	db := ToDB(c)
	if err := store.InsertAccount(db, na); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, &M{"account": na})
}

// func CreateAccount(c *echo.Context) error {
// 	u := &m.Account{}
// 	if err := c.Bind(u); err != nil {
// 		return c.JSON(500, err)
// 	}
// 	log.Println(u)
//
// 	u, err := m.NewAccountFor(u.AccountID)
// 	if err != nil {
// 		c.Error(err)
// 		return err
// 	}
// 	log.Println(u)
//
// 	db := ToDB(c)
// 	if err := db.Create(u).Error; err != nil {
// 		return c.JSON(503, err)
// 	}
//
// 	// c.JSON(201, &APIRes{"user": []m.Account{*newAccount}})
// 	return c.JSON(201, u)
// }

// ActivateAccount will activate a user account from a token or an admin.
func ActivateAccount(c *echo.Context) error {
	return c.JSON(http.StatusNotImplemented, "activate user not implemented")
}
