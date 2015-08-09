package store

import (
	"testing"

	_ "github.com/lib/pq"
)

// func getDB() *sqlx.DB {
// 	db := sqlx.MustConnect("postgres", "postgres://localhost/gogrademe-api-test?sslmode=disable")
// 	db.MapperFunc(snaker.CamelToSnake)
//
// 	db.MustExec(`DROP DATABASE IF EXISTS "gogrademe-api-test"`)
// 	db.MustExec(`CREATE DATABASE "gogrademe-api-test" TEMPLATE "gogrademe-api-tmpl"`)
// 	return db
// }
//
// func TestPersonInsert(t *testing.T) {
// 	db := getDB()
// 	defer db.Close()
// 	person := &model.Person{FirstName: "fred"}
// 	if err := InsertPerson(db, person); err != nil {
// 		t.Error(err)
// 	}
//
// 	fmt.Println(person)
// }
//
// func TestPersonGet(t *testing.T) {
// 	db := getDB()
// 	defer db.Close()
// 	m, err := GetPersonList(db)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	t.Logf("%+v", m)
// }

func TestUpdate(t *testing.T) {
	expected := "UPDATE account SET first_name=:first_name, last_name=:last_name, middle_name=:middle_name WHERE id = :id"
	stmt := Update("account").SetN("first_name", "last_name").SetN("middle_name").Eq("id").String()
	if stmt != expected {
		t.Errorf("expected: %s got: %s", expected, stmt)
	}
}
