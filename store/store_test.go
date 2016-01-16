package store

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestUpdate(t *testing.T) {
	expected := "UPDATE account SET first_name=:first_name, last_name=:last_name, middle_name=:middle_name WHERE id = :id"
	stmt := Update("account").SetN("first_name", "last_name").SetN("middle_name").Eq("id").String()
	if stmt != expected {
		t.Errorf("expected: %s got: %s", expected, stmt)
	}
}
