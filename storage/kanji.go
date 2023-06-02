package storage

import (
	"context"
	"fmt"
)

const selectKanijById = `
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username, kanji.meanings, kanji.example_sentences, kanji.example_words
	FROM kanji
	WHERE kanji.id = $1
		;`

type Kanji struct {
	ID               string `json:"id,omitempty"`
	ExampleSentences string `json:"exampleSentences,omitempty"`
	ExampleWords     string `json:"exampleWords,omitempty"`
	Kanji            string `json:"kanji,omitempty"`
	KanjiRating      int    `json:"kanjiRating,omitempty"`
	KunReading       string `json:"kunReading,omitempty"`
	Meanings         string `json:"meanings,omitempty"`
	OnReading        string `json:"onReading,omitempty"`
	Username         string `json:"username,omitempty"`
}

type CreateKanjiRequest struct {
	ExampleSentences string
	ExampleWords     string
	Kanji            string
	KanjiRating      int
	KunReading       string
	Meanings         string
	OnReading        string
	Username         string
}

type UpdateKanjiRequest struct {
	ID               *string `json:"id,omitempty"`
	Kanji            *string `json:"kanji,omitempty"`
	ExampleSentences *string `json:"exampleSentences,omitempty"`
	ExampleWords     *string `json:"exampleWords,omitempty"`
	OnReading        *string `json:"onReading,omitempty"`
	KunReading       *string `json:"kunReading,omitempty"`
	KanjiRating      *int    `json:"kanjiRating,omitempty"`
	Meanings         *string `json:"meanings,omitempty"`
}

func ScanKanji(s Scanner) (*Kanji, error) {
	k := &Kanji{}
	if err := s.Scan(&k.ID, &k.Kanji, &k.OnReading, &k.KunReading, &k.KanjiRating, &k.Username, &k.Meanings, &k.ExampleSentences, &k.ExampleWords); err != nil {
		return nil, err
	}

	return k, nil
}

func (s *Storage) CreateKanji(ctx context.Context, k CreateKanjiRequest) (*Kanji, error) {
	// open transaction
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	insertKanjiStatement := "INSERT INTO kanji(kanji, on_reading, kun_reading, kanji_rating, username, meanings, example_sentences, example_words) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id, kanji, on_reading, kun_reading, kanji_rating, username, meanings, example_sentences, example_words;"
	insertedKanjiRow := tx.QueryRowContext(ctx, insertKanjiStatement, k.Kanji, k.OnReading, k.KunReading, k.KanjiRating, k.Username, k.Meanings, k.ExampleSentences, k.ExampleWords)

	kanji, err := ScanKanji(insertedKanjiRow)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return kanji, nil
}

func (s *Storage) GetKanji(ctx context.Context, id string) (*Kanji, error) {
	row := s.conn.QueryRowContext(ctx, selectKanijById, id)
	return ScanKanji(row)
}

func (s *Storage) GetAllKanji(ctx context.Context) ([]*Kanji, error) {
	selectStatement := `
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username,kanji.meanings, kanji.example_sentences, kanji.example_words
	FROM kanji
	`
	rows, err := s.conn.QueryContext(ctx, selectStatement)
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
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// update kanji
	updateKanjiStatement := "UPDATE kanji SET kanji = COALESCE($1, kanji), on_reading = COALESCE($2, on_reading), kun_reading = COALESCE($3, kun_reading), kanji_rating = COALESCE($4, kanji_rating), meanings = COALESCE($5, meanings), example_sentences = COALESCE($6, example_sentences), example_words = COALESCE($7, example_words) WHERE id = $8 RETURNING id"
	_, err = tx.ExecContext(ctx, updateKanjiStatement, k.Kanji, k.OnReading, k.KunReading, k.KanjiRating, k.Meanings, k.ExampleSentences, k.ExampleWords, k.ID)
	if err != nil {
		return nil, err
	}

	// select from kanji, kanji_sentence, and kanji_example
	row := tx.QueryRowContext(ctx, selectKanijById, k.ID)
	kanji, scanErr := ScanKanji(row)

	// commit transaction
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return kanji, scanErr

}

func (s *Storage) DeleteKanji(id string) error {
	_, err := s.conn.Exec("DELETE FROM kanji WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
