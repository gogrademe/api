package store

import "github.com/gogrademe/api/model"

// GetAssignmentGroup --
func (s *Store) GetAssignmentGroup(id int) (*model.AssignmentGroup, error) {
	var r model.AssignmentGroup
	return &r, s.db.Get(&r, "select * from assignment_group WHERE id=$1", id)
}

// GetAssignmentGroupList --
func (s *Store) GetAssignmentGroupList() ([]model.AssignmentGroup, error) {
	var r []model.AssignmentGroup
	return r, s.db.Select(&r, "select * from assignment_group")
}

// InsertAssignmentGroup --
func (s *Store) InsertAssignmentGroup(assignmentGroup *model.AssignmentGroup) error {
	stmt := `INSERT INTO assignment_group (name, weight, course_id, term_id, created_at, updated_at)
			 VALUES (:name, :weight,:course_id, :term_id, :due_date, :created_at, :updated_at) RETURNING id`
	assignmentGroup.UpdateTime()

	var err error
	assignmentGroup.ID, err = insert(s.db, stmt, assignmentGroup)
	return err
}

// Update --
func (s *Store) UpdateAssignmentGroup(assignmentGroup *model.AssignmentGroup) error {
	stmt := Update("assignment_group").SetN("name", "weight", "course_id", "term_id", "created_at", "updated_at").Eq("id").String()
	assignmentGroup.UpdateTime()

	_, err := s.db.NamedQuery(stmt, assignmentGroup)
	return err

}

// Del --
func (s *Store) DeleteAssignmentGroup(id int) error {
	stmt := `DELETE FROM assignment_group WHERE id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
