package handler

import (
	m "github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

// GetAllUsers http endpoint to return all users.
func GetAllUsers(c *echo.Context) error {
	db := ToDB(c)

	users := &[]m.User{}
	db.Find(&users)

	return c.JSON(200, users)

}

func CreateUser(c *echo.Context) error {

	u := &m.User{}
	if err := c.Bind(u); err != nil {
		return err
	}

	// newUser, err := m.NewUserFor(u.Email, u.PersonID)

	db := ToDB(c)
	if err := db.Create(u).Error; err != nil {
		return c.JSON(503, err)
	}

	// c.JSON(201, &APIRes{"user": []m.User{*newUser}})
	return c.JSON(201, u)
}
