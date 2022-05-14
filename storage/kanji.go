package storage

import (
	"context"
	"fmt"
)

type Kanji struct {
	ID          string
	Kanji       string
	OnReading   string
	KunReading  string
	KanjiRating int
	Username    string
}

type CreateKanjiRequest struct {
	Kanji       string
	OnReading   string
	KunReading  string
	KanjiRating int
	Username    string
}

type UpdateKanjiRequest struct {
	ID          *string `json:"id,omitempty"`
	Kanji       *string `json:"kanji,omitempty"`
	OnReading   *string `json:"onReading,omitempty"`
	KunReading  *string `json:"kunReading,omitempty"`
	KanjiRating *int    `json:"kanjiRating,omitempty"`
}

func ScanKanji(s Scanner) (*Kanji, error) {
	k := &Kanji{}
	if err := s.Scan(&k.ID, &k.Kanji, &k.OnReading, &k.KunReading, &k.KanjiRating, &k.Username); err != nil {
		return nil, err
	}

	return k, nil
}

func (s *Storage) CreateKanji(ctx context.Context, k CreateKanjiRequest) (*Kanji, error) {
	insertStatement := "INSERT INTO kanji(kanji, on_reading, kun_reading, kanji_rating, username)  VALUES($1,$2,$3,$4,$5) RETURNING id, kanji,on_reading, kun_reading, kanji_rating, username"
	row := s.conn.QueryRowContext(ctx, insertStatement, k.Kanji, k.OnReading, k.KunReading, k.KanjiRating, k.Username)
	return ScanKanji(row)
}

func (s *Storage) GetKanji(ctx context.Context, id string) (*Kanji, error) {
	selectStatement := "SELECT * FROM kanji WHERE id = $1"
	row := s.conn.QueryRowContext(ctx, selectStatement, id)
	return ScanKanji(row)
}

func (s *Storage) GetAllKanji(ctx context.Context) ([]*Kanji, error) {
	rows, err := s.conn.QueryContext(ctx, "SELECT id, kanji, on_reading, kun_reading, kanji_rating, username FROM kanji")
	if err != nil {
		return nil, fmt.Errorf("could not retrieve items: %w", err)
	}
	defer rows.Close()

	var kanjiSlice []*Kanji
	for rows.Next() {
		kanji, err := ScanKanji(rows)
		if err != nil {
			return nil, fmt.Errorf("could not scan item: %w", err)
		}

		kanjiSlice = append(kanjiSlice, kanji)
	}

	return kanjiSlice, nil
}

func (s *Storage) UpdateKanji(ctx context.Context, k UpdateKanjiRequest) (*Kanji, error) {
	updateStatement := "UPDATE kanji SET kanji = COALESCE($1, kanji), on_reading = COALESCE($2, on_reading), kun_reading = COALESCE($3, kun_reading), kanji_rating = COALESCE($4, kanji_rating) WHERE id = $5 RETURNING id, kanji, on_reading, kun_reading, kanji_rating, username"
	row := s.conn.QueryRowContext(ctx, updateStatement, k.Kanji, k.OnReading, k.KunReading, k.KanjiRating, k.ID)
	return ScanKanji(row)
}

func (s *Storage) DeleteKanji(id string) error {
	_, err := s.conn.Exec("DELETE FROM kanji WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
