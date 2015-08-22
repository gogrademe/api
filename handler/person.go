package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

func CreatePerson(c *echo.Context) error {
	p := &model.Person{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.InsertPerson(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}

func GetAllPeople(c *echo.Context) error {
	db := ToDB(c)

	ppl, err := db.GetPersonList()
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func GetPerson(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))
	ppl, err := db.GetPerson(id)
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func DeletePerson(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DeletePerson(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func UpdatePerson(c *echo.Context) error {
	p := &model.Person{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.UpdatePerson(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}
