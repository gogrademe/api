package store

import "github.com/gogrademe/api/model"

// GetCourse --
func (s *Store) GetCourse(id int) (*model.Course, error) {
	var r model.Course
	err := s.ru.SelectDoc("course.*", "level.name AS gradelevel").
		Many("terms", `SELECT * from term WHERE id IN (SELECT id FROM course_term WHERE course_id=$1)`, id).
		From(`course
			INNER JOIN level ON (course.level_id = level.id)`).
		Where("course.id = $1", id).QueryStruct(&r)

	return &r, err
}

// GetCourseList --
func (s *Store) GetCourseList() ([]model.Course, error) {
	var r []model.Course
	return r, s.db.Select(&r, "select course.*, level.name AS grade_level FROM course INNER JOIN level ON (course.level_id = level.id)")
}

// InsertCourse --
func (s *Store) InsertCourse(course *model.Course) error {
	stmt := `INSERT INTO course (name, level_id, max_students, created_at, updated_at)
			 VALUES (:name, :level_id, :max_students, :created_at, :updated_at) RETURNING id`
	course.UpdateTime()

	var err error
	course.ID, err = insert(s.db, stmt, course)
	return err
}

func (s *Store) InsertCourseTerm(courseID, termID int) error {

	_, err := s.ru.InsertInto("course_term").Columns("course_id", "term_id").Values(courseID, termID).Exec()
	return err
}

// Update --
func (s *Store) UpdateCourse(course *model.Course) error {
	stmt := Update("course").SetN("name", "level_id", "max_students", "created_at", "updated_at").Eq("id").String()
	course.UpdateTime()

	_, err := s.db.NamedQuery(stmt, course)
	return err

}

// Del --
func (s *Store) DeleteCourse(id int) error {
	stmt := `DELETE FROM course WHERE id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
