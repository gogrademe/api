package store

import (
	"fmt"
	"testing"

	"github.com/gogrademe/api/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/serenize/snaker"
)

func getDB() *sqlx.DB {
	db := sqlx.MustConnect("postgres", "postgres://localhost/gogrademe-api-test?sslmode=disable")
	db.MapperFunc(snaker.CamelToSnake)

	db.MustExec(`DROP DATABASE IF EXISTS "gogrademe-api-test"`)
	db.MustExec(`CREATE DATABASE "gogrademe-api-test" TEMPLATE "gogrademe-api-tmpl"`)
	return db
}

func TestPersonInsert(t *testing.T) {
	db := getDB()
	defer db.Close()
	person := &model.Person{FirstName: "fred"}
	if err := InsertPerson(db, person); err != nil {
		t.Error(err)
	}

	fmt.Println(person)
}

func TestPersonGet(t *testing.T) {
	db := getDB()
	defer db.Close()
	m, err := GetPersonList(db)
	if err != nil {
		t.Error(err)
	}

	t.Logf("%+v", m)
}
