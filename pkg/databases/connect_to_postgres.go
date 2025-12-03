package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToPostgres(host string, port int, database string, user string, password string) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		host, port, database, user, password,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf(
			"connecting to the database %s on %s:%d via user %s: %w",
			database, host, port, user, err,
		)
	}

	return db, nil
}
