package adapters

import (
	"database/sql"
	"mon-backend/domain/kanji"
	"mon-backend/domain/reading"
)

type repositories struct {
	ReadingRepository reading.ReadingRepository
	KanjiRepository   kanji.KanjiRepository
}

func InitRepositories(db *sql.DB) repositories {
	readingRepository := NewReadingRepository(db)
	kanjiRepository := NewKanjiRepository(db)
	return repositories{ReadingRepository: readingRepository, KanjiRepository: kanjiRepository}
}
