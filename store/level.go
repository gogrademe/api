package store

import "github.com/gogrademe/api/model"

// GetLevel --
func (s *Store) GetLevel(id int) (*model.Level, error) {
	var r model.Level
	return &r, s.db.Get(&r, "select * from level WHERE id=$1", id)
}

// GetLevelList --
func (s *Store) GetLevelList() ([]model.Level, error) {
	var r []model.Level
	return r, s.db.Select(&r, "select * from level")
}

// InsertLevel --
func (s *Store) InsertLevel(level *model.Level) error {
	stmt := `INSERT INTO level (name,school_year, created_at, updated_at)
			 VALUES (:name, :school_year, :created_at, :updated_at) RETURNING id`
	level.UpdateTime()

	var err error
	level.ID, err = insert(s.db, stmt, level)
	return err
}

// Update --
func (s *Store) UpdateLevel(level *model.Level) error {
	stmt := Update("level").SetN("name", "school_year", "created_at", "updated_at").Eq("id").String()
	level.UpdateTime()

	_, err := s.db.NamedQuery(stmt, level)
	return err

}

// Del --
func (s *Store) DeleteLevel(id int) error {
	stmt := `DELETE FROM level WHERE id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
