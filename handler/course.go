package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

func CreateCourse(c *echo.Context) error {
	p := &model.Course{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.InsertCourse(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, p)
}

func CreateCourseTerm(c *echo.Context) error {
	courseID, _ := strconv.Atoi(c.Param("courseID"))
	termID, _ := strconv.Atoi(c.Param("termID"))

	db := ToDB(c)
	if err := db.InsertCourseTerm(courseID, termID); err != nil {
		return ErrSaving.Log(err)
	}
	return c.NoContent(http.StatusCreated)
}

func GetAllCourses(c *echo.Context) error {
	db := ToDB(c)

	ppl, err := db.GetCourseList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, ppl)

}

func GetCourse(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))
	ppl, err := db.GetCourse(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, ppl)

}

func DeleteCourse(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DeleteCourse(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func UpdateCourse(c *echo.Context) error {
	p := &model.Course{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.UpdateCourse(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, p)
}
