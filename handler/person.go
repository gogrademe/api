package handler

import (
	"net/http"
	"strconv"

	"github.com/gogrademe/api/model"
	"github.com/gogrademe/api/store"
	"github.com/labstack/echo"
)

func CreatePerson(c *echo.Context) error {
	p := &model.Person{}
	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	db := ToDB(c)
	if err := store.InsertPerson(db, p); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, &M{"person": p})
}

func GetAllPeople(c *echo.Context) error {
	db := ToDB(c)

	ppl, err := store.GetPersonList(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, &M{"person": ppl})

}

func GetPerson(c *echo.Context) error {
	db := ToDB(c)
	id, _ := strconv.Atoi(c.Param("id"))
	ppl, err := store.GetPerson(db, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, &M{"person": ppl})

}

//
// // UpdatePerson ...
// func UpdatePerson(c *gin.Context) {
//
// 	id := c.Param("id")
//
// 	p.ID = id
// 	err := store.People.Update(p, id)
//
// 	if err != nil {
// 		writeError(c.Writer, "Error updating Person", 500, err)
// 		return
// 	}
//
// 	c.JSON(200, &APIRes{"person": []m.Person{*p}})
// 	return
// }
//

//
// // DeletePerson ...
// func DeletePerson(c *gin.Context) {
//
// 	id := c.Param("id")
//
// 	_, err := store.DB.RunWrite(store.People.Get(id).Delete())
// 	if err == store.ErrNotFound {
// 		writeError(c.Writer, notFoundError, 404, nil)
// 		return
// 	}
// 	if err != nil {
// 		writeError(c.Writer, serverError, 500, nil)
// 		return
// 	}
//
// 	c.JSON(200, &APIRes{"person": []m.Person{}})
// 	return
// }
