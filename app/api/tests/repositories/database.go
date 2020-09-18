package repositories

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

func InitialiseTestDatabase(t *testing.T, url string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", url)
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}
