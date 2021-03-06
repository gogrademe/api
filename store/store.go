package store

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"

	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"

	"github.com/gogrademe/api/model"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/serenize/snaker"
)

type Filter map[string]interface{}

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

func (s *Store) EnsureAdmin() {

	accts, err := s.GetAccountList()
	if err != nil {
		log.Fatal("error retrieving accounts: ", err)
	}

	if len(accts) > 0 {
		return
	}

	p := &model.Person{
		FirstName: "admin",
	}

	if err := s.InsertPerson(p); err != nil {
		log.Fatal(err)
	}

	u, _ := model.NewAccountFor(p.PersonID, "admin@host.local")
	if err := u.SetPassword("admin123"); err != nil {
		log.Fatal(err)
	}

	u.SetActive()

	if err := s.InsertAccount(u); err != nil {
		log.Fatal(err)
	}

	log.Println("created default admin account")

}

func Connect(addr string) *Store {
	// db := sqlx.MustConnect("pgx", addr)
	db := sqlx.MustConnect("postgres", addr)
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)
	db.MapperFunc(snaker.CamelToSnake)
	dat.EnableInterpolation = true

	return &Store{db: db, ru: runner.NewDB(db.DB, "postgres")}
}
