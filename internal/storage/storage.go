package storage

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/oisinmulvihill/go-webdev/internal/core"
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

func AddUser(conn *pgx.Conn, username string, password string) error {
	query := `INSERT INTO users (username, password) VALUES ('bob', 'bobpassword')`
	if _, err := conn.Exec(context.Background(), query); err != nil {
		return err
	}

	return nil
}

func GetUsers(conn *pgx.Conn) ([]core.User, error) {

	rows, err := conn.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []core.User
	for rows.Next() {

		var u core.User
		var createdAt time.Time
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &createdAt)
		if err != nil {
			return nil, err
		}
		u.CreatedAt = createdAt.Format(time.RFC3339)
		users = append(users, u)
	}

	return users, nil
}
