package handler

import (
	"strconv"

	"github.com/labstack/echo"
)

func GetGradebook(c *echo.Context) error {
	db := ToDB(c)
	course, _ := strconv.Atoi(c.Param("courseID"))
	term, _ := strconv.Atoi(c.Param("termID"))
	res, err := db.GetCourseTermAttemptList(course, term)
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, res)

}
