package store

import "github.com/gogrademe/api/model"

// InsertSession --
func (s *Store) InsertSession(session *model.Session) error {
	stmt := `INSERT INTO session (token, account_id, expires_at, created_at, updated_at)
			 VALUES (:token, :account_id, :expires_at, :created_at, :updated_at) RETURNING id`
	session.UpdateTime()

	var err error
	session.ID, err = insert(s.db, stmt, session)
	return err
}
