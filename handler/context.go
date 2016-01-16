package handler

import (
	"github.com/gogrademe/api/model"
	"github.com/gogrademe/api/store"
	"github.com/labstack/echo"
)

func SetDB(db *store.Store) echo.HandlerFunc {
	return func(c *echo.Context) error {
		c.Set("db", db)
		return nil
	}
}

func ToDB(ctx *echo.Context) *store.Store {
	return ctx.Get("db").(*store.Store)
}

func ToClaims(ctx *echo.Context) map[string]interface{} {
	return ctx.Get("claims").(map[string]interface{})
}

func ToMe(ctx *echo.Context) *model.Account {
	// ok, id := ToClaims(ctx)["account_id"]
	// if
	return nil
}

// return &model.Session{AccountID: claims["account_id"]}
