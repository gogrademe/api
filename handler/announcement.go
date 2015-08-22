package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

func CreateAnnouncement(c *echo.Context) error {
	p := &model.Announcement{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.InsertAnnouncement(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}

func GetAllAnnouncements(c *echo.Context) error {
	db := ToDB(c)

	ppl, err := db.GetAnnouncementList()
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func GetAnnouncement(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))
	ppl, err := db.GetAnnouncement(id)
	if err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(200, ppl)

}

func DeleteAnnouncement(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := db.DeleteAnnouncement(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func UpdateAnnouncement(c *echo.Context) error {
	p := &model.Announcement{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	db := ToDB(c)
	if err := db.UpdateAnnouncement(p); err != nil {
		return ErrServerError.Log(err)
	}

	return c.JSON(http.StatusCreated, p)
}
