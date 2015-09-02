package store

import "github.com/gogrademe/api/model"

// GetLevel --
func (s *Store) GetLevel(id int) (*model.Level, error) {
	var r model.Level
	return &r, s.db.Get(&r, "select * from level WHERE level_id=$1", id)
}

// GetLevelList --
func (s *Store) GetLevelList() ([]model.Level, error) {
	var r []model.Level
	return r, s.db.Select(&r, "select * from level")
}

// InsertLevel --
func (s *Store) InsertLevel(level *model.Level) error {
	stmt := `INSERT INTO level (name, created_at, updated_at)
			 VALUES (:name, :created_at, :updated_at) RETURNING level_id`

	var err error
	level.LevelID, err = insert(s.db, stmt, level)
	return err
}

// Update --
func (s *Store) UpdateLevel(level *model.Level) error {
	stmt := Update("level").SetN("name", "created_at", "updated_at").Eq("level_id").String()

	_, err := s.db.NamedQuery(stmt, level)
	return err

}

// Del --
func (s *Store) DeleteLevel(id int) error {
	stmt := `DELETE FROM level WHERE level_id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
