package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

func CreateLevel(c echo.Context) error {
	p := &model.Level{}
	if err := c.Bind(p); err != nil {
		return ErrBind.Log(err)
	}

	db := ToDB(c)
	if err := db.InsertLevel(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}

func GetAllLevels(c echo.Context) error {
	db := ToDB(c)

	ppl, err := db.GetLevelList()
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func GetLevel(c echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))
	ppl, err := db.GetLevel(id)
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func DeleteLevel(c echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DeleteLevel(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func UpdateLevel(c echo.Context) error {
	p := &model.Level{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.UpdateLevel(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}
