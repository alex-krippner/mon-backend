package adapters

import (
	"database/sql"
	"mon-backend/domain/kanji"
	"mon-backend/domain/reading"
	"mon-backend/domain/vocabulary"
)

type repositories struct {
	ReadingRepository reading.ReadingRepository
	KanjiRepository   kanji.KanjiRepository
	VocabRepository   vocabulary.VocabRepository
}

func InitRepositories(db *sql.DB) repositories {
	readingRepository := NewReadingRepository(db)
	kanjiRepository := NewKanjiRepository(db)
	vocabRepository := NewVocabRepository(db)
	return repositories{ReadingRepository: readingRepository, KanjiRepository: kanjiRepository, VocabRepository: vocabRepository}
}
