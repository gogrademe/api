package store

import "github.com/gogrademe/api/model"

// GetAssignment --
func (s *Store) GetAssignment(id int) (*model.Assignment, error) {
	var r model.Assignment
	return &r, s.db.Get(&r, `SELECT assignment.*, row_to_json(assignment_group) as group FROM assignment
		INNER JOIN assignment_group on(assignment_id)
		WHERE assignment_id=$1", id`)
}

// GetAssignmentList --
func (s *Store) GetAssignmentList() ([]model.Assignment, error) {
	var r []model.Assignment
	stmt := `SELECT assignment.*,
	 		 assignment_group.name as "group.name",
			 assignment_group.weight as "group.weight"
			 FROM assignment INNER JOIN assignment_group USING(group_id)
			 ORDER BY due_date, "group.name", assignment.name`
	return r, s.db.Select(&r, stmt)
}

// GetAssignmentList --
// func (s *Store) GetAssignmentListOld() ([]model.Assignment, error) {
// 	var r []model.Assignment
// 	return r, s.ru.SelectDoc(`assignment.assignment_id,
// 		assignment.name,
// 		max_score,
// 		assignment.due_date::timestamptz,
// 		assignment.created_at::timestamptz,
// 		assignment.updated_at::timestamptz`).
// 		One("group", `SELECT name, weight, course_id, term_id FROM assignment_group WHERE group_id = assignment.group_id`).
// 		From(`assignment`).QueryStructs(&r)
// 	// return r, s.db.Select(&r, `SELECT assignment.*, row_to_json(assignment_group.*) "group" FROM assignment INNER JOIN assignment_group USING(group_id)`)
// }

// GetAssignmentList --
func (s *Store) GetCourseAssignmentList(courseID, termID int) ([]model.Assignment, error) {
	var r []model.Assignment
	return r, s.db.Select(&r, "SELECT a.*, row_to_json(g.*) as group FROM assignment a INNER JOIN assignment_group g ON g.id = a.group_id WHERE a.course_id = $1 AND a.term_id = $2", courseID, termID)
}

// InsertAssignment --
func (s *Store) InsertAssignment(assignment *model.Assignment) error {
	stmt := `INSERT INTO assignment (name, course_id, term_id, group_id, max_score, due_date, created_at, updated_at)
			 VALUES (:name, :course_id, :term_id, :group_id, :max_score, :due_date, :created_at, :updated_at) RETURNING assignment_id`

	var err error
	assignment.AssignmentID, err = insert(s.db, stmt, assignment)
	return err
}

// Update --
func (s *Store) UpdateAssignment(assignment *model.Assignment) error {
	stmt := Update("assignment").SetN("name", "course_id", "term_id", "group_id", "max_score", "due_date", "created_at", "updated_at").Eq("assignment_id").String()

	_, err := s.db.NamedQuery(stmt, assignment)
	return err

}

// Del --
func (s *Store) DeleteAssignment(id int) error {
	stmt := `DELETE FROM assignment WHERE assignment_id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
