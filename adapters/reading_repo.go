package adapters

import (
	"context"
	"database/sql"
	"fmt"

	"mon-backend/domain/reading"
)

func ScanReading(s Scanner) (*reading.Reading, error) {
	r := &reading.Reading{}
	if err := s.Scan(&r.ID, &r.Translation, &r.Japanese, &r.Title, &r.Username); err != nil {
		return nil, err
	}

	return r, nil
}

type ReadingRepository struct {
	db *sql.DB
}

func NewReadingRepository(db *sql.DB) *ReadingRepository {
	return &ReadingRepository{
		db: db,
	}
}

func (r ReadingRepository) CreateReading(ctx context.Context, req *reading.Reading) (*reading.Reading, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	insertReadingStatement := "INSERT INTO reading(translation, japanese, title, username) VALUES($1, $2, $3, $4) RETURNING id, translation, japanese, title, username;"
	insertedReadingRow := tx.QueryRowContext(ctx, insertReadingStatement, req.Translation, req.Japanese, req.Title, req.Username)
	reading, err := ScanReading(insertedReadingRow)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return reading, nil

}

func (r ReadingRepository) GetAllReading(ctx context.Context, username string) ([]*reading.Reading, error) {

	selectStatement := "SELECT reading.id, reading.translation, reading.japanese, reading.title, reading.username FROM reading WHERE reading.username = $1"

	rows, err := r.db.QueryContext(ctx, selectStatement, username)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve readings: %w", err)
	}

	defer rows.Close()

	var readingSlice []*reading.Reading
	for rows.Next() {
		reading, err := ScanReading(rows)
		if err != nil {
			return nil, fmt.Errorf("cloud not scan reading: %w", err)
		}

		readingSlice = append(readingSlice, reading)
	}

	return readingSlice, nil
}

func (r ReadingRepository) UpdateReading(ctx context.Context, reading reading.Reading) (*reading.Reading, error) {
	selectReadingById := "SELECT reading.id, reading.translation, reading.japanese, reading.title, reading.username FROM reading WHERE reading.id = $1"

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	updateReadingStatement := "UPDATE reading SET translation = COALESCE($1, translation), japanese = COALESCE($2, japanese), title = COALESCE($3, title), username = COALESCE($4, username) WHERE id = $5 RETURNING id"
	_, err = tx.ExecContext(ctx, updateReadingStatement, reading.Translation, reading.Japanese, reading.Title, reading.Username, reading.ID)
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(ctx, selectReadingById, reading.ID)
	updatedReading, scanErr := ScanReading(row)
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return updatedReading, scanErr
}

func (r ReadingRepository) DeleteReading(id string) error {
	_, err := r.db.Exec("DELETE FROM reading WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
