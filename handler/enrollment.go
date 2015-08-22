package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

func CreateEnrollment(c *echo.Context) error {
	p := &model.Enrollment{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.InsertEnrollment(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}

func GetAllEnrollments(c *echo.Context) error {
	db := ToDB(c)

	ppl, err := db.GetEnrollmentList()
	if err != nil {

		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func GetEnrollment(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))
	ppl, err := db.GetEnrollment(id)
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func DeleteEnrollment(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DeleteEnrollment(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func UpdateEnrollment(c *echo.Context) error {
	p := &model.Enrollment{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.UpdateEnrollment(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}
