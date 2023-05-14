package storage

import (
	"context"
	"fmt"
)

type Reading struct {
	ID          string `json:"id,omitempty"`
	Translation string `json:"translation,omitempty"`
	Japanese    string `json:"japanese,omitempty"`
	Title       string `json:"title,omitempty"`
}

type CreateReadingRequest struct {
	Translation string
	Japanese    string
	Title       string
}

func ScanReading(s Scanner) (*Reading, error) {
	r := &Reading{}
	if err := s.Scan(&r.ID, &r.Translation, &r.Japanese, &r.Title); err != nil {
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

	insertReadingStatement := "INSERT INTO reading(translation, japanese, title) VALUES($1, $2, $3) RETURNING id, translation, japanese, title;"
	insertedReadingRow := tx.QueryRowContext(ctx, insertReadingStatement, r.Translation, r.Japanese, r.Title)
	reading, err := ScanReading(insertedReadingRow)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return reading, nil

}

func (s *Storage) GetAllReading(ctx context.Context) ([]*Reading, error) {
	selectStatement := "SELECT reading.id, reading.translation, reading.japanese, reading.title FROM reading"

	rows, err := s.conn.QueryContext(ctx, selectStatement)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve readings: %w", err)
	}

	defer rows.Close()

	var readingSlice []*Reading
	for rows.Next() {
		reading, err := ScanReading(rows)
		if err != nil {
			return nil, fmt.Errorf("cloud not scan reading: %w", err)
		}

		readingSlice = append(readingSlice, reading)
	}
	return readingSlice, nil
}

func (s *Storage) UpdateReading(ctx context.Context, r Reading) (*Reading, error) {
	selectReadingById := "SELECT reading.id, reading.translation, reading.japanese, reading.title FROM reading WHERE reading.id = $1"

	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	updateReadingStatement := "UPDATE reading SET translation = COALESCE($1, translation), japanese = COALESCE($2, japanese), title = COALESCE($3, title) WHERE id = $4 RETURNING id"
	_, err = tx.ExecContext(ctx, updateReadingStatement, r.Translation, r.Japanese, r.Title, r.ID)
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(ctx, selectReadingById, r.ID)
	reading, scanErr := ScanReading(row)

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return reading, scanErr
}

func (s *Storage) DeleteReading(id string) error {
	_, err := s.conn.Exec("DELETE FROM reading WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
