package store

import "github.com/gogrademe/api/model"

// GetAttempt --
func (s *Store) GetAttempt(id int) (*model.Attempt, error) {
	var r model.Attempt
	return &r, s.db.Get(&r, "select * from attempt WHERE attempt_id=$1", id)
}

// GetAttemptList --
// func (s *Store) GetAttemptList() ([]model.AttemptResponse, error) {
// 	var r []model.AttemptResponse
// 	stmt := `SELECT attempt.*,
// 	person.display_name as "student.display_name",
// 	assignment.name as "assignment.name"
// 	FROM attempt
// 	INNER JOIN person USING(person_id)
// 	INNER JOIN assignment USING(assignment_id)`
// 	return r, s.db.Select(&r, stmt)
// }

// GetAttemptList --
func (s *Store) GetAttemptList() ([]model.Attempt, error) {
	var r []model.Attempt
	stmt := `select * from attempt`

	// return r, s.db.Select(&r, stmt)
	// s.db.QueryRow(stmt).Scan(&r[0])
	return r, s.db.Select(&r, stmt)
}

// InsertAttempt --
func (s *Store) InsertAttempt(attempt *model.Attempt) error {
	stmt := `INSERT INTO attempt (score, average, person_id, assignment_id, created_at, updated_at)
			 VALUES (:score, :average, :person_id, :assignment_id, :created_at, :updated_at) RETURNING attempt_id`

	var err error
	attempt.AttemptID, err = insert(s.db, stmt, attempt)
	return err
}

// Update --
func (s *Store) UpdateAttempt(attempt *model.Attempt) error {
	stmt := Update("attempt").SetN("score", "average", "assignment_id", "person_id", "created_at", "updated_at").Eq("id").String()

	_, err := s.db.NamedQuery(stmt, attempt)
	return err

}

// Del --
func (s *Store) DeleteAttempt(id int) error {
	stmt := `DELETE FROM attempt WHERE id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
