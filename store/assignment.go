package store

import "github.com/gogrademe/api/model"

// GetAssignment --
func (s *Store) GetAssignment(id int) (*model.Assignment, error) {
	var r model.Assignment
	return &r, s.db.Get(&r, "select * from assignment WHERE id=$1", id)
}

// GetAssignmentList --
func (s *Store) GetAssignmentList() ([]model.Assignment, error) {
	var r []model.Assignment
	return r, s.db.Select(&r, "select * from assignment")
}

// InsertAssignment --
func (s *Store) InsertAssignment(assignment *model.Assignment) error {
	stmt := `INSERT INTO assignment (name, course_id, term_id, group_id, max_score, due_date, created_at, updated_at)
			 VALUES (:name, :course_id, :term_id, :group_id, :max_score, :due_date, :created_at, :updated_at) RETURNING id`
	assignment.UpdateTime()

	var err error
	assignment.ID, err = insert(s.db, stmt, assignment)
	return err
}

// Update --
func (s *Store) UpdateAssignment(assignment *model.Assignment) error {
	stmt := Update("assignment").SetN("name", "course_id", "term_id", "group_id", "max_score", "due_date", "created_at", "updated_at").Eq("id").String()
	assignment.UpdateTime()

	_, err := s.db.NamedQuery(stmt, assignment)
	return err

}

// Del --
func (s *Store) DeleteAssignment(id int) error {
	stmt := `DELETE FROM assignment WHERE id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
