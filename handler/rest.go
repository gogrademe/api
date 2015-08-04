package handler

import (
	"net/http"

	m "github.com/gogrademe/api/model"
	"github.com/labstack/echo"
)

// GetAllType --
func GetAllType(c *echo.Context) error {
	db := ToDB(c)

	var typ []m.Type
	if err := db.Select(&typ, "SELECT name FROM type"); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, &M{"type": typ})

}

//GetType --
func GetType(c *echo.Context) error {
	id, db := c.Param("id"), ToDB(c)

	var typ m.Type
	if err := db.Get(&typ, "SELECT name FROM type WHERE id = $1", id); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, &M{"type": typ})

}

// CreateType --
// func CreateType(c *echo.Context) error {
// 	typ := &m.Type{}
// 	if err := c.Bind(typ); err != nil {
// 		return c.JSON(500, err)
// 	}
//
// 	db := ToDB(c)
// 	if err := db.Create(typ).Error; err != nil {
// 		return c.JSON(503, err)
// 	}
//
// 	// c.JSON(201, &APIRes{"user": []m.Type{*newType}})
// 	return c.JSON(201, &M{"type": typ})
// }
