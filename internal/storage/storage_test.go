package storage

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/peterldowns/pgtestdb"
)

func NewTestDB(t *testing.T) *pgx.Conn {
	conf := pgtestdb.Config{
		DriverName: "pgx",
		User:       "postgres",
		Password:   "password",
		Host:       "localhost",
		Port:       "7432",
		Options:    "sslmode=disable",
	}
	var migrator pgtestdb.Migrator = pgtestdb.NoopMigrator{}
	dbConf := pgtestdb.Custom(t, conf, migrator)

	conn, err := Connection(dbConf.URL())
	if err != nil {
		t.Fatal(err)
	}

	setUpError := Init(conn)
	if setUpError != nil {
		t.Fatal(setUpError)
	}

	return conn
}

func TestTotalUsersEmptyDB(t *testing.T) {
	t.Parallel()

	conn := NewTestDB(t)
	defer conn.Close(context.Background())

	userCount, err := TotalUsers(conn)
	if err != nil {
		t.Fatal(err)
	}
	if userCount != 0 {
		t.Errorf("expected 0 users but got %d", userCount)
	}

}

func TestTotalUsersWithUsersPresent(t *testing.T) {
	t.Parallel()

	conn := NewTestDB(t)
	defer conn.Close(context.Background())

	userCount, err := TotalUsers(conn)
	if err != nil {
		t.Fatal(err)
	}
	if userCount != 0 {
		t.Errorf("expected 0 users but got %d", userCount)
	}

	query := `INSERT INTO users (username, password) VALUES ('bob', 'bobpassword')`
	if _, err := conn.Exec(context.Background(), query); err != nil {
		t.Fatalf("Failed to insert a user: %s", err)
	}

	userCount, err = TotalUsers(conn)
	if err != nil {
		t.Fatal(err)
	}
	if userCount != 1 {
		t.Errorf("expected 1 users but got %d", userCount)
	}
}

func TestAddUser(t *testing.T) {

	t.Parallel()

	conn := NewTestDB(t)
	defer conn.Close(context.Background())

	// no user should be present:
	var total int
	err := conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		t.Fatal(err)
	}
	if total != 0 {
		t.Errorf("expected 0 users but got %d", total)
	}

	// add a user
	AddUser(conn, "bob", "bobpassword")

	err = conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM users where username = 'bob'").Scan(&total)
	if err != nil {
		t.Fatal(err)
	}
	if total != 1 {
		t.Errorf("expected 1 users but got %d", total)
	}
}
