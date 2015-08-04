package store

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

const insertAnnouncement = `:title, :person_id, :posted_date`
const insertAssignment = `:name, :class_id, :term_id, :group_id, :max_score, :due_date`
const insertCourse = `:name, :grade_level, :max_students`
const insertEnrollment = `:person_id, :class_id, :term_id`
const insertPerson = `:first_name, :middle_name, :last_name, :grade_level`
const insertContactInfo = `:phone_number, :email`
const insertSession = `:token, :user_id, :expires_at`
const insertAccount = `:person_id, :email, :role, :hashed_password, :activation_token, :disabled`

func insert(db *sqlx.DB, stmt string, params interface{}) (int, error) {
	rows, err := db.NamedQuery(stmt, params)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		err := rows.Scan(&id)
		return id, err
	}

	return 0, errors.New("No serial value returned for insert: " + stmt + ", error: " + rows.Err().Error())
}
