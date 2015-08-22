package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

func CreateAssignmentGroup(c *echo.Context) error {
	p := &model.AssignmentGroup{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.InsertAssignmentGroup(p); err != nil {
		return ErrSaving.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}

func GetAllAssignmentGroups(c *echo.Context) error {
	db := ToDB(c)

	ppl, err := db.GetAssignmentGroupList()
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func GetAssignmentGroup(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))
	ppl, err := db.GetAssignmentGroup(id)
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func DeleteAssignmentGroup(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DeleteAssignmentGroup(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func UpdateAssignmentGroup(c *echo.Context) error {
	p := &model.AssignmentGroup{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.UpdateAssignmentGroup(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}
