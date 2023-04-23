package storage

import (
	"context"
)

type Reading struct {
	ID                 string `json:"id,omitempty"`
	EnglishTranslation string `json:"englishTranslation,omitempty"`
	Japanese           string `json:"japanese,omitempty"`
}

type CreateReadingRequest struct {
	EnglishTranslation string
	Japanese           string
}

func ScanReading(s Scanner) (*Reading, error) {
	r := &Reading{}
	if err := s.Scan(&r.ID, &r.EnglishTranslation, &r.Japanese); err != nil {
		return nil, err
	}

	return r, nil
}

func (s *Storage) CreateReading(ctx context.Context, r CreateReadingRequest) (*Reading, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	insertReadingStatement := "INSERT INTO reading(english_translation, japanese) VALUES($1, $2) RETURNING id, english_translation, japanese;"
	insertedReadingRow := tx.QueryRowContext(ctx, insertReadingStatement, r.EnglishTranslation, r.Japanese)
	reading, err := ScanReading(insertedReadingRow)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return reading, nil

}
