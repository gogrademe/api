package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

func SetDB(db *sqlx.DB) echo.HandlerFunc {
	return func(c *echo.Context) error {
		c.Set("db", db)
		return nil
	}
}

func ToDB(c *echo.Context) *sqlx.DB {
	return c.Get("db").(*sqlx.DB)
}
