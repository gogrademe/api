package store

import "github.com/gogrademe/api/model"

// GetEnrollment --
func (s *Store) GetEnrollment(id int) (*model.Enrollment, error) {
	var r model.Enrollment
	return &r, s.db.Get(&r, "select * from enrollment WHERE enrollment_id=$1", id)
}

// GetEnrollmentList --
func (s *Store) GetEnrollmentList() ([]model.Enrollment, error) {
	var r []model.Enrollment
	stmt := `SELECT enrollment.*,
	 				person.first_name as "person.first_name",
					person.middle_name as "person.middle_name",
					person.last_name as "person.last_name",
					person.display_name as "person.display_name",
					person.grade_level as "person.grade_level"
				from enrollment INNER JOIN person using(person_id)`
	return r, s.db.Select(&r, stmt)
}

// InsertEnrollment --
func (s *Store) InsertEnrollment(enrollment *model.Enrollment) error {
	stmt := `INSERT INTO enrollment (person_id, course_id, term_id, created_at, updated_at)
			 VALUES (:person_id, :course_id, :term_id, :created_at, :updated_at) RETURNING enrollment_id`

	var err error
	enrollment.EnrollmentID, err = insert(s.db, stmt, enrollment)
	return err
}

// Update --
func (s *Store) UpdateEnrollment(enrollment *model.Enrollment) error {
	stmt := Update("enrollment").SetN("person_id", "course_id", "term_id", "created_at", "updated_at").Eq("enrollment_id").String()

	_, err := s.db.NamedQuery(stmt, enrollment)
	return err

}

// Del --
func (s *Store) DeleteEnrollment(id int) error {
	stmt := `DELETE FROM enrollment WHERE enrollment_id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
