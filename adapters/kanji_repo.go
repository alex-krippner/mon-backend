package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"mon-backend/domain/kanji"
)

const selectKanijById = `
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username, kanji.meanings, kanji.example_sentences, kanji.example_words
	FROM kanji
	WHERE kanji.id = $1
	`

func ScanKanji(s Scanner) (*kanji.Kanji, error) {
	k := &kanji.Kanji{}
	if err := s.Scan(&k.ID, &k.Kanji, &k.OnReading, &k.KunReading, &k.KanjiRating, &k.Username, &k.Meanings, &k.ExampleSentences, &k.ExampleWords); err != nil {
		return nil, err
	}

	return k, nil
}

type KanjiRepository struct {
	db *sql.DB
}

func NewKanjiRepository(db *sql.DB) *KanjiRepository {
	return &KanjiRepository{
		db: db,
	}
}

func (r KanjiRepository) AddKanji(ctx context.Context, req *kanji.Kanji) (*kanji.Kanji, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	insertKanjiStatement := "INSERT INTO kanji(kanji, on_reading, kun_reading, kanji_rating, username, meanings, example_sentences, example_words) VALUES($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id, kanji, on_reading, kun_reading, kanji_rating, username, meanings, example_sentences, example_words;"
	insertedKanjiRow := tx.QueryRowContext(ctx, insertKanjiStatement, req.Kanji, req.OnReading, req.KunReading, req.KanjiRating, req.Username, req.Meanings, req.ExampleSentences, req.ExampleWords)

	kanji, err := ScanKanji(insertedKanjiRow)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return kanji, nil
}

func (r KanjiRepository) GetAllKanji(ctx context.Context, username string) ([]*kanji.Kanji, error) {

	selectStatement := `
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username, kanji.meanings, kanji.example_sentences, kanji.example_words
	FROM kanji WHERE kanji.username = $1
	`
	rows, err := r.db.QueryContext(ctx, selectStatement, username)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve items: %w", err)
	}
	defer rows.Close()

	var kanjiSlice []*kanji.Kanji
	for rows.Next() {
		kanji, err := ScanKanji(rows)
		if err != nil {
			return nil, fmt.Errorf("could not scan item: %w", err)
		}

		kanjiSlice = append(kanjiSlice, kanji)
	}
	return kanjiSlice, nil
}

func (r KanjiRepository) GetKanji(ctx context.Context, id string) (*kanji.Kanji, error) {
	row := r.db.QueryRowContext(ctx, selectKanijById, id)
	return ScanKanji(row)
}

func (r KanjiRepository) UpdateKanji(ctx context.Context, k kanji.Kanji) (*kanji.Kanji, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// update kanji
	updateKanjiStatement := "UPDATE kanji SET kanji = COALESCE($1, kanji), on_reading = COALESCE($2, on_reading), kun_reading = COALESCE($3, kun_reading), kanji_rating = COALESCE($4, kanji_rating), meanings = COALESCE($5, meanings), example_sentences = COALESCE($6, example_sentences), example_words = COALESCE($7, example_words), username = COALESCE($8, username) WHERE id = $9 RETURNING id"
	_, err = tx.ExecContext(ctx, updateKanjiStatement, k.Kanji, k.OnReading, k.KunReading, k.KanjiRating, k.Meanings, k.ExampleSentences, k.ExampleWords, k.Username, k.ID)
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

func (r KanjiRepository) DeleteKanji(id string) error {
	_, err := r.db.Exec("DELETE FROM kanji WHERE id = $1", id)

	if err != nil {
		return err
	}

	return nil
}
