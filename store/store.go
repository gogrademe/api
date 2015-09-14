package store

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"

	"github.com/jmoiron/sqlx"
	"github.com/serenize/snaker"
)

type Store struct {
	db *sqlx.DB
	ru *runner.DB
}

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

// func (s *Store) UpdateAccount(account *model.Account) error {
// 	stmt := `UPDATE account SET
// 				person_id = :person_id,
// 				email = :email,
// 				role = :role,
// 				hashed_password = :hashed_password,
// 				activation_token = :activation_token,
// 				disabled = :disabled,
// 				created_at = :created_at,
// 				updated_at = :updated_at
// 			WHERE id = :id`
// 	account.UpdateTime()
//
// 	_, err := s.db.NamedQuery(stmt, account)
// 	return err
// }
type UpdateStmt struct {
	table  string
	values []value
	where  string
}

type value struct {
	col, opr, name string
}

func (v value) String() string {
	return fmt.Sprint(v.col, v.opr, v.name)
}

func Update(table string) *UpdateStmt {
	return &UpdateStmt{table: table}
}

func (stmt *UpdateStmt) SetN(name ...string) *UpdateStmt {
	for k := range name {
		stmt.values = append(stmt.values, value{col: name[k], opr: "=", name: ":" + name[k]})
	}
	return stmt
}

func (stmt *UpdateStmt) Eq(name string) *UpdateStmt {
	stmt.where = fmt.Sprintf("%[1]s = :%[1]s", name)
	return stmt
}

func (stmt *UpdateStmt) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("UPDATE %s SET ", stmt.table))

	vals := make([]string, len(stmt.values))
	for i := range stmt.values {
		vals[i] = stmt.values[i].String()
	}

	buf.WriteString(strings.Join(vals, ", "))
	buf.WriteString(fmt.Sprintf(" WHERE %s", stmt.where))

	return buf.String()
}
func Connect(addr string) *Store {
	dburi, _ := url.Parse(addr)
	db := sqlx.MustConnect(dburi.Scheme, dburi.String())
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	dat.EnableInterpolation = true

	db.MapperFunc(snaker.CamelToSnake)
	return &Store{db: db, ru: runner.NewDB(db.DB, "postgres")}
}
