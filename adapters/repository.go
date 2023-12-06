package adapters

import (
	"database/sql"
	"mon-backend/domain/reading"
)

type repositories struct {
	ReadingRepository reading.ReadingRepository
}

func InitRepositories(db *sql.DB) repositories {
	readingRepository := NewReadingRepository(db)
	return repositories{ReadingRepository: readingRepository}
}
