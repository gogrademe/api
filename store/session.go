package store

import (
	"github.com/gogrademe/api/model"
	"github.com/jmoiron/sqlx"
)

// InsertSession --
func InsertSession(db *sqlx.DB, session *model.Session) error {
	stmt := `INSERT INTO session (:token, :user_id, :expires_at, created_at, updated_at)
			 VALUES (:token, :user_id, :expires_at, :created_at, :updated_at) RETURNING id`
	session.UpdateTime()

	var err error
	session.ID, err = insert(db, stmt, session)
	return err
}
