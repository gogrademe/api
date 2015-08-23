package store

import "github.com/gogrademe/api/model"

// GetTerm --
func (s *Store) GetTerm(id int) (*model.Term, error) {
	var r model.Term
	return &r, s.db.Get(&r, "select * from term WHERE term_id=$1", id)
}

// GetTermList --
func (s *Store) GetTermList() ([]model.Term, error) {
	var r []model.Term
	return r, s.db.Select(&r, "select * from term")
}

// InsertTerm --
func (s *Store) InsertTerm(term *model.Term) error {
	stmt := `INSERT INTO term (name,school_year, created_at, updated_at)
			 VALUES (:name, :school_year, :created_at, :updated_at) RETURNING term_id`
	term.UpdateTime()

	var err error
	term.TermID, err = insert(s.db, stmt, term)
	return err
}

// Update --
func (s *Store) UpdateTerm(term *model.Term) error {
	stmt := Update("term").SetN("name", "school_year", "created_at", "updated_at").Eq("term_id").String()
	term.UpdateTime()

	_, err := s.db.NamedQuery(stmt, term)
	return err

}

// Del --
func (s *Store) DeleteTerm(id int) error {
	stmt := `DELETE FROM term WHERE term_id=$1`

	_, err := s.db.Exec(stmt, id)
	return err

}
