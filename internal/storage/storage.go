package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connection(databaseDSN string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), databaseDSN)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	return conn, nil
}

func Init(conn *pgx.Conn) error {
	// plain text password is terrible. It should be a strong hash of the
	// password. Plaintext is what the tutorial used and I'll stick with it
	// for now :(
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		log.Printf("Unabled to create users table because: %v", err)
		return err
	}

	log.Printf("Tables created OK if they did not exist already.")

	return nil
}

func TotalUsers(conn *pgx.Conn) (int, error) {
	var total int

	err := conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM users").Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
