package handler

import (
	"github.com/gogrademe/api/store"
	"github.com/labstack/echo"
)

func SetDB(db *store.Store) echo.HandlerFunc {
	return func(c *echo.Context) error {
		c.Set("db", db)
		return nil
	}
}

func ToDB(c *echo.Context) *store.Store {
	return c.Get("db").(*store.Store)
}
