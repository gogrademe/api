package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

func SetDB(db *gorm.DB) echo.HandlerFunc {
	return func(c *echo.Context) error {
		c.Set("db", db)
		return nil
	}
}

func ToDB(c *echo.Context) *gorm.DB {
	return c.Get("db").(*gorm.DB)
}
