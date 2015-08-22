package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

func CreateAssignment(c *echo.Context) error {
	p := &model.Assignment{}
	if err := c.Bind(p); err != nil {
		return ErrBind.Log(err)
	}

	db := ToDB(c)
	if err := db.InsertAssignment(p); err != nil {
		return ErrSaving.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}

func GetAllAssignments(c *echo.Context) error {
	db := ToDB(c)

	ppl, err := db.GetAssignmentList()
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func GetCourseAssignments(c *echo.Context) error {
	db := ToDB(c)
	course, _ := strconv.Atoi(c.Param("courseID"))
	term, _ := strconv.Atoi(c.Param("termID"))
	ppl, err := db.GetCourseAssignmentList(course, term)
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func GetAssignment(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))
	ppl, err := db.GetAssignment(id)
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func DeleteAssignment(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DeleteAssignment(id); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusOK, nil)
}

func UpdateAssignment(c *echo.Context) error {
	p := &model.Assignment{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.UpdateAssignment(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}