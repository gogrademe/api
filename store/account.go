package store

import "github.com/gogrademe/api/model"

// GetAccount --
func (s *Store) GetAccount(id int) (*model.Account, error) {
	var r model.Account
	return &r, s.db.Get(&r, "select * from account WHERE account_id=$1", id)
}

// GetAccountEmail --
func (s *Store) GetAccountEmail(email string) (*model.Account, error) {
	var r model.Account
	return &r, s.db.Get(&r, "select * from account WHERE email=$1", email)
}

// GetAccountByToken --
func (s *Store) GetAccountByToken(token string) (*model.Account, error) {
	var r model.Account
	return &r, s.db.Get(&r, "select * from account WHERE activation_token=$1", token)
}

// GetAccountList --
func (s *Store) GetAccountList() ([]model.Account, error) {
	var r []model.Account
	return r, s.db.Select(&r, "select * from account")
}

// InsertAccount --
func (s *Store) InsertAccount(account *model.Account) error {
	stmt := `INSERT INTO account (person_id, email, hashed_password, activation_token, disabled, created_at, updated_at)
			 VALUES (:person_id, :email, :hashed_password, :activation_token, :disabled, :created_at, :updated_at) RETURNING account_id`
	var err error
	account.AccountID, err = insert(s.db, stmt, account)
	return err
}

// Update --
func (s *Store) UpdateAccount(account *model.Account) error {
	stmt := Update("account").SetN(
		"person_id",
		"email",
		"hashed_password",
		"activation_token",
		"disabled",
		"created_at",
		"updated_at").Eq("id").String()

	_, err := s.db.NamedQuery(stmt, account)
	return err
}

// Del --
func (s *Store) DeleteAccount(id int) error {
	stmt := `DELETE FROM account WHERE id=$1`

	_, err := s.db.Exec(stmt, id)
	return err
}
