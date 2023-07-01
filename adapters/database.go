package adapters

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	apiServerStorageDatabaseURL string = "DATABASE_URL"
)

type Scanner interface {
	Scan(dest ...interface{}) error
}

func GetDatabase() (*sql.DB, error) {

	databaseURL := os.Getenv(apiServerStorageDatabaseURL)
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("could not open sql: %w", err)
	}

	return db, nil
}
