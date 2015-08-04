package store

import (
	"github.com/gogrademe/api/model"
	"github.com/jmoiron/sqlx"
)

// GetAccount --
func GetAccount(db *sqlx.DB, id int) (*model.Account, error) {
	var r model.Account
	return &r, db.Get(&r, "select * from account WHERE id=$1", id)
}

// GetAccountEmail --
func GetAccountEmail(db *sqlx.DB, email string) (*model.Account, error) {
	var r model.Account
	return &r, db.Get(&r, "select * from account WHERE email=$1", email)
}

// GetAccountList --
func GetAccountList(db *sqlx.DB) ([]model.Account, error) {
	var r []model.Account
	return r, db.Select(&r, "select * from account")
}

// InsertAccount --
func InsertAccount(db *sqlx.DB, account *model.Account) error {
	// _ := []string{"person_id", "email", "role", "hashed_password", "activation_token", "disabled", "created_at", "updated_at"}
	stmt := `INSERT INTO account (person_id, email, role, hashed_password, activation_token, disabled, created_at, updated_at)
			 VALUES (:person_id, :email, :role, :hashed_password, :activation_token, :disabled, :created_at, :updated_at) RETURNING id`
	account.UpdateTime()

	var err error
	account.ID, err = insert(db, stmt, account)
	return err
}

// Update --
func UpdateAccount(db *sqlx.DB, account *model.Account) error {
	return nil
}

// Del --
func DelAccount(db *sqlx.DB, account *model.Account) error {
	return nil
}
