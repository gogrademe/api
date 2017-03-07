package store

import (
	"github.com/Sirupsen/logrus"
	"github.com/gogrademe/api/model"
)

// GetCourse --
func (s *Store) GetCourse(id int) (*model.Course, error) {
	var r model.Course
	stmt := `SELECT course.*, level.name as grade_level
			FROM course INNER JOIN level USING(level_id)
			WHERE course.course_id = $1`
	return &r, s.db.Get(&r, stmt, id)
}

// GetCourseList --
func (s *Store) GetCourseList() ([]model.Course, error) {
	var r []model.Course
	stmt := `SELECT course.*, level.name AS grade_level
	FROM course INNER JOIN level USING(level_id)`
	return r, s.db.Select(&r, stmt)
}

// InsertCourse --
func (s *Store) InsertCourse(course *model.Course) error {
	// stmt := `WITH course as (
	// 		INSERT INTO course (name, level_id, max_students, created_at, updated_at)
	// 		VALUES (:name, :level_id, :max_students, :created_at, :updated_at) RETURNING course_id
	// 		)
	// 		INSERT INTO course_term (course_id, term_id)
	// 		VALUES (course_id,:term_id)`
	// var err error
	// course.CourseID, err = insert(s.db, stmt, course)

	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}

	rows, err := tx.NamedQuery(`INSERT INTO course (name, level_id, max_students, created_at, updated_at)
	VALUES (:name, :level_id, :max_students, :created_at, :updated_at) RETURNING course_id`, course)
	if err != nil {
		return err
	}

	var id int
	if rows.Next() {
		err = rows.Scan(&id)
	}
	if err != nil {
		return err
	}
	rows.Close()

	logrus.Info(id, course.Terms[0].TermID)
	_, err = tx.Exec("INSERT INTO course_term (course_id, term_id) VALUES ($1, $2)", id, course.Terms[0].TermID)
	if err != nil {
		logrus.Info(err)
		return err
	}

	// return 0, errors.New("No serial value returned for insert: " + stmt + ", error: " + rows.Err().Error())
	return tx.Commit()
}

func (s *Store) InsertCourseTerm(courseID, termID int) error {
	_, err := s.ru.InsertInto("course_term").Columns("course_id", "term_id").Values(courseID, termID).Exec()
	return err
}

// Update --
func (s *Store) UpdateCourse(course *model.Course) error {
	stmt := Update("course").SetN("name", "level_id", "max_students", "created_at", "updated_at").Eq("course_id").String()
	_, err := s.db.NamedQuery(stmt, course)
	return err
}

// Del --
func (s *Store) DeleteCourse(id int) error {
	stmt := `DELETE FROM course WHERE course_id=$1`
	_, err := s.db.Exec(stmt, id)
	return err
}
