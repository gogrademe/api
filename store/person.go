package store

import (
	"github.com/gogrademe/api/model"
	"github.com/jmoiron/sqlx"
)

// GetPerson --
func GetPerson(db *sqlx.DB, id int) (*model.Person, error) {
	var r model.Person
	return &r, db.Get(&r, "select * from person WHERE id=$1", id)
}

// GetPersonList --
func GetPersonList(db *sqlx.DB) ([]model.Person, error) {
	var r []model.Person
	return r, db.Select(&r, "select * from person")
}

// InsertPerson --
func InsertPerson(db *sqlx.DB, person *model.Person) error {
	stmt := `INSERT INTO person (first_name, middle_name, last_name, grade_level, created_at, updated_at)
			 VALUES (:first_name, :middle_name, :last_name, :grade_level, :created_at, :updated_at) RETURNING id`
	person.UpdateTime()

	var err error
	person.ID, err = insert(db, stmt, person)
	return err
}

// Update --
func Update(db *sqlx.DB, person *model.Person) error {
	return nil
}

// Del --
func Del(db *sqlx.DB, person *model.Person) error {
	return nil
}
