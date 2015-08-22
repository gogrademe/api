package store

import "github.com/gogrademe/api/model"

// GetEnrollment --
func (s *Store) GetEnrollment(id int) (*model.Enrollment, error) {
	var r model.Enrollment
	return &r, s.db.Get(&r, "select * from enrollment WHERE id=$1", id)
}

// GetEnrollmentList --
func (s *Store) GetEnrollmentList() ([]model.Enrollment, error) {
	var r []model.Enrollment
	// err := s.ru.SelectDoc("enrollment.*").
	// 	One("person", `SELECT * from person WHERE person.id = enrollment.person_id`).
	// 	From("enrollment").
	// 	QueryStructs(&r)
	q := `SELECT *, (SELECT row_to_json(dat__person.*) FROM (SELECT * FROM person WHERE person.id = enrollment.person_id) AS dat__person) AS person FROM enrollment`
	return r, s.db.Select(&r, q)
}

// InsertEnrollment --
func (s *Store) InsertEnrollment(enrollment *model.Enrollment) error {
	stmt := `INSERT INTO enrollment (person_id, course_id, term_id, created_at, updated_at)
			 VALUES (:person_id, :course_id, :term_id, :created_at, :updated_at) RETURNING id`
	enrollment.UpdateTime()

	var err error
	enrollment.ID, err = insert(s.db, stmt, enrollment)
	return err
}

// Update --
func (s *Store) UpdateEnrollment(enrollment *model.Enrollment) error {
	stmt := Update("enrollment").SetN("person_id", "course_id", "term_id", "created_at", "updated_at").Eq("id").String()
	enrollment.UpdateTime()

	_, err := s.db.NamedQuery(stmt, enrollment)
	return err

}

// Del --
func (s *Store) DeleteEnrollment(id int) error {
	stmt := `DELETE FROM enrollment WHERE id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}