package storage

import (
	"database/sql"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/peterldowns/pgtestdb"
)

func NewTestDB(t *testing.T) (*sql.DB, *pgx.Conn) {
	conf := pgtestdb.Config{
		DriverName: "pgx",
		User:       "postgres",
		Password:   "password",
		// Database:   fmt.Sprintf("webdev-test-%d", rand.Intn(99999)),
		Host:    "localhost",
		Port:    "7432",
		Options: "sslmode=disable",
	}
	var migrator pgtestdb.Migrator = pgtestdb.NoopMigrator{}
	db := pgtestdb.New(t, conf, migrator)
	defer db.Close()

	conn, err := Connection(conf.URL())
	if err != nil {
		t.Fatal(err)
	}

	setUpError := Init(conn)
	if setUpError != nil {
		t.Fatal(setUpError)
	}

	return db, conn
}

func TestTotalUsersEmptyDB(t *testing.T) {
	t.Parallel()

	_, conn := NewTestDB(t)

	userCount, err := TotalUsers(conn)
	if err != nil {
		t.Fatal(err)
	}
	if userCount != 0 {
		t.Errorf("expected 0 users but got %d", userCount)
	}

}
