package store

import "github.com/gogrademe/api/model"

// GetAssignmentGroup --
func (s *Store) GetAssignmentGroup(id int) (*model.AssignmentGroup, error) {
	var r model.AssignmentGroup
	return &r, s.db.Get(&r, "select * from assignment_group WHERE group_id=$1", id)
}

// GetAssignmentGroupList --
func (s *Store) GetAssignmentGroupList(courseID, termID int) ([]model.AssignmentGroup, error) {
	var r []model.AssignmentGroup
	return r, s.db.Select(&r, "select * from assignment_group where course_id=$1 and term_id=$2", courseID, termID)
}

// InsertAssignmentGroup --
func (s *Store) InsertAssignmentGroup(assignmentGroup *model.AssignmentGroup) error {
	stmt := `INSERT INTO assignment_group (name, weight, course_id, term_id, created_at, updated_at)
			 VALUES (:name, :weight, :course_id, :term_id, :created_at, :updated_at) RETURNING group_id`

	var err error
	assignmentGroup.GroupID, err = insert(s.db, stmt, assignmentGroup)
	return err
}

// Update --
func (s *Store) UpdateAssignmentGroup(assignmentGroup *model.AssignmentGroup) error {
	stmt := Update("assignment_group").SetN("name", "weight", "course_id", "term_id", "created_at", "updated_at").Eq("group_id").String()

	_, err := s.db.NamedQuery(stmt, assignmentGroup)
	return err

}

// Del --
func (s *Store) DeleteAssignmentGroup(id int) error {
	stmt := `DELETE FROM assignment_group WHERE group_id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
