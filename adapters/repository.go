package adapters

import (
	"database/sql"
	"mon-backend/domain"
)

type repositories struct {
	ReadingRepository domain.ReadingRepository
}

func InitRepositories(db *sql.DB) repositories {
	readingRepository := NewReadingRepository(db)
	return repositories{ReadingRepository: readingRepository}
}
