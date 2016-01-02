package store

import "github.com/gogrademe/api/model"

// GetPerson --
func (s *Store) GetPerson(id int) (*model.Person, error) {
	var r model.Person
	return &r, s.db.Get(&r, "select *, person.display_name from person WHERE person_id=$1", id)
}

// GetPersonList --
func (s *Store) GetPersonList() ([]model.Person, error) {
	var r []model.Person
	return r, s.db.Select(&r, "select *, person.display_name from person")
}

// InsertPerson --
func (s *Store) InsertPerson(person *model.Person) error {
	stmt := `INSERT INTO person (first_name, middle_name, last_name, grade_level , created_at, updated_at)
			 VALUES (:first_name, :middle_name, :last_name, :grade_level, :created_at, :updated_at) RETURNING person_id`

	var err error
	person.PersonID, err = insert(s.db, stmt, person)
	return err
}

// Update --
func (s *Store) UpdatePerson(person *model.Person) error {
	stmt := Update("person").SetN("first_name", "middle_name", "last_name", "grade_level", "created_at", "updated_at").Eq("person_id").String()

	_, err := s.db.NamedQuery(stmt, person)
	return err

}

// Del --
func (s *Store) DeletePerson(id int) error {
	stmt := `DELETE FROM person WHERE person_id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
