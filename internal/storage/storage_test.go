package storage

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/peterldowns/pgtestdb"
)

func testingConnectionHelper(t *testing.T) *pgx.Conn {
	conf := pgtestdb.Config{
		DriverName: "pgx",
		User:       "service",
		Password:   "service",
		Host:       "127.0.0.1",
		Port:       "7432",
		Options:    "sslmode=disable",
	}
	var migrator pgtestdb.Migrator = pgtestdb.NoopMigrator{}
	db := pgtestdb.New(t, conf, migrator)
	defer db.Close()

	conn, err := Connection(conf.URL())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close(context.Background())

	setUpError := Init(conn)
	if setUpError != nil {
		t.Fatal(setUpError)
	}

	return conn
}

func TestTotalUsersEmptyDB(t *testing.T) {
	t.Parallel()

	conn := testingConnectionHelper(t)

	userCount, err := TotalUsers(conn)
	if err != nil {
		t.Fatal(err)
	}
	if userCount != 0 {
		t.Errorf("expected 0 users but got %d", userCount)
	}

}
