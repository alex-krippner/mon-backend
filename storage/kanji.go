package storage

import (
	"context"
	"fmt"
)

type Kanji struct {
	ID              string
	Kanji           string
	ExampleSentence string
	ExampleWord     string
	OnReading       string
	KunReading      string
	KanjiRating     int
	Username        string
}

type CreateKanjiRequest struct {
	Kanji           string
	ExampleSentence string
	ExampleWord     string
	OnReading       string
	KunReading      string
	KanjiRating     int
	Username        string
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
	if err := s.Scan(&k.ID, &k.Kanji, &k.OnReading, &k.KunReading, &k.KanjiRating, &k.Username, &k.ExampleSentence, &k.ExampleWord); err != nil {
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
	// defer rollback
	defer tx.Rollback()
	// insert into kanji
	insertKanjiStatement := "INSERT INTO kanji(kanji, on_reading, kun_reading, kanji_rating, username)  VALUES($1,$2,$3,$4,$5);"
	_, err = tx.ExecContext(ctx, insertKanjiStatement, k.Kanji, k.OnReading, k.KunReading, k.KanjiRating, k.Username)
	if err != nil {
		return nil, err
	}
	// insert into kanji_sentence
	insertKanjiSentenceStatement := "INSERT INTO kanji_sentence(kanji_id, example_sentence) VALUES((SELECT id FROM kanji WHERE kanji.kanji = $1 AND kanji.username = $3), $2);"
	_, err = tx.ExecContext(ctx, insertKanjiSentenceStatement, k.Kanji, k.ExampleSentence, k.Username)
	if err != nil {
		return nil, err
	}
	fmt.Println(k.ExampleWord)
	// insert into kanji_example
	insertKanjiExampleStatement := "INSERT INTO kanji_example(kanji_id, example_word) VALUES((SELECT id FROM kanji WHERE kanji.kanji = $1 AND kanji.username = $3), $2);"
	_, err = tx.ExecContext(ctx, insertKanjiExampleStatement, k.Kanji, k.ExampleWord, k.Username)
	if err != nil {
		return nil, err
	}

	// select from kanji and kanji_sentenece
	selectKanijStatement := `
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username, kanji_sentence.example_sentence, kanji_example.example_word 
	FROM kanji
		JOIN kanji_sentence ON kanji.id = kanji_sentence.kanji_id  
		JOIN kanji_example ON kanji.id = kanji_example.kanji_id
	WHERE kanji.kanji = $1 AND kanji.username = $2;`

	row := tx.QueryRowContext(ctx, selectKanijStatement, k.Kanji, k.Username)
	kanji, scanErr := ScanKanji(row)
	// commit transaction
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return kanji, scanErr
}

func (s *Storage) GetKanji(ctx context.Context, id string) (*Kanji, error) {
	selectStatement := "SELECT * FROM kanji WHERE id = $1"
	row := s.conn.QueryRowContext(ctx, selectStatement, id)
	return ScanKanji(row)
}

func (s *Storage) GetAllKanji(ctx context.Context) ([]*Kanji, error) {
	rows, err := s.conn.QueryContext(ctx, "SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating,kanji.username, kanji_sentence.example_sentence FROM kanji JOIN kanji_sentence ON kanji.id = kanji_sentence.kanji_id")
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
