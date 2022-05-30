package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

const selectKanijById = `
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username, ks.example_sentences, ke.example_words
	FROM kanji
		LEFT JOIN (
			SELECT ks.kanji_id AS id, json_agg(
				json_build_object('exampleSentence', example_sentence, 'sentenceId', sentence_id, 'kanjiId', kanji_id)) AS example_sentences
			FROM kanji_sentence ks
			GROUP BY ks.kanji_id
		) ks USING(id)
		LEFT JOIN (
			SELECT ke.kanji_id AS id, json_agg(
				json_build_object('exampleWord', example_word, 'wordId', word_id, 'kanjiId', kanji_id)) AS example_words
			FROM kanji_example ke
			GROUP BY ke.kanji_id
		) ke USING (id)
	WHERE kanji.id = $1
		;`

type ExampleSentence struct {
	SentenceID      string `json:"sentenceId,omitempty"`
	KanjiID         string `json:"kanjiId,omitempty"`
	ExampleSentence string `json:"exampleSentence,omitempty"`
}

type ExampleWord struct {
	WordId      string `json:"wordId,omitempty"`
	KanjiID     string `json:"kanjiId,omitempty"`
	ExampleWord string `json:"exampleWord,omitempty"`
}
type ExampleSentences []ExampleSentence
type ExampleWords []ExampleWord

type Kanji struct {
	ID               string           `json:"id,omitempty"`
	Kanji            string           `json:"kanji,omitempty"`
	ExampleSentences ExampleSentences `json:"exampleSentences,omitempty"`
	ExampleWords     ExampleWords     `json:"exampleWords,omitempty"`
	OnReading        string           `json:"onReading,omitempty"`
	KunReading       string           `json:"kunReading,omitempty"`
	KanjiRating      int              `json:"kanjiRating,omitempty"`
	Username         string           `json:"username,omitempty"`
}

type CreateKanjiRequest struct {
	Kanji            string
	ExampleSentences ExampleSentences
	ExampleWords     ExampleWords
	OnReading        string
	KunReading       string
	KanjiRating      int
	Username         string
}

type UpdateKanjiRequest struct {
	ID               *string          `json:"id,omitempty"`
	Kanji            *string          `json:"kanji,omitempty"`
	ExampleSentences ExampleSentences `json:"exampleSentences,omitempty"`
	ExampleWords     ExampleWords     `json:"exampleWords,omitempty"`
	OnReading        *string          `json:"onReading,omitempty"`
	KunReading       *string          `json:"kunReading,omitempty"`
	KanjiRating      *int             `json:"kanjiRating,omitempty"`
}

//TODO: Remove function duplication. Try composition, embedding, or factory
func (e *ExampleSentences) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")

	}
	return json.Unmarshal(b, &e)
}

func (e *ExampleWords) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")

	}
	return json.Unmarshal(b, &e)
}

func ScanKanji(s Scanner) (*Kanji, error) {
	k := &Kanji{}
	if err := s.Scan(&k.ID, &k.Kanji, &k.OnReading, &k.KunReading, &k.KanjiRating, &k.Username, &k.ExampleSentences, &k.ExampleWords); err != nil {
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
	insertKanjiStatement := "INSERT INTO kanji(kanji, on_reading, kun_reading, kanji_rating, username)  VALUES($1,$2,$3,$4,$5) RETURNING id;"
	insertedKanjiRow := tx.QueryRowContext(ctx, insertKanjiStatement, k.Kanji, k.OnReading, k.KunReading, k.KanjiRating, k.Username)
	var kanjiId string
	insertedKanjiRow.Scan(&kanjiId)
	if err != nil {
		return nil, err
	}

	// insert into kanji_sentence
	insertKanjiSentenceStatement := "INSERT INTO kanji_sentence(kanji_id, example_sentence) VALUES($1, $2);"
	for _, exampleSentence := range k.ExampleSentences {
		_, err = tx.ExecContext(ctx, insertKanjiSentenceStatement, kanjiId, exampleSentence.ExampleSentence)
		if err != nil {
			return nil, err
		}
	}

	// insert into kanji_example
	insertKanjiExampleStatement := "INSERT INTO kanji_example(kanji_id, example_word) VALUES($1, $2);"
	for _, exampleWord := range k.ExampleWords {
		_, err = tx.ExecContext(ctx, insertKanjiExampleStatement, kanjiId, exampleWord.ExampleWord)
		if err != nil {
			return nil, err
		}
	}
	// select from kanji, kanji_sentence, and kanji_example
	row := tx.QueryRowContext(ctx, selectKanijById, kanjiId)
	kanji, scanErr := ScanKanji(row)
	// commit transaction
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return kanji, scanErr
}

func (s *Storage) GetKanji(ctx context.Context, id string) (*Kanji, error) {
	row := s.conn.QueryRowContext(ctx, selectKanijById, id)
	return ScanKanji(row)
}

func (s *Storage) GetAllKanji(ctx context.Context) ([]*Kanji, error) {
	selectStatement := `
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username, ks.example_sentences, ke.example_words
	FROM kanji
		LEFT JOIN (
			SELECT ks.kanji_id AS id, json_agg(
				json_build_object('exampleSentence', example_sentence, 'sentenceId', sentence_id, 'kanjiId', kanji_id)) AS example_sentences
			FROM kanji_sentence ks
			GROUP BY ks.kanji_id
		) ks USING(id)
		LEFT JOIN (
			SELECT ke.kanji_id AS id, json_agg(
				json_build_object('exampleWord', example_word, 'wordId', word_id, 'kanjiId', kanji_id)) AS example_words
			FROM kanji_example ke
			GROUP BY ke.kanji_id
		) ke USING (id);
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
	updateKanjiStatement := "UPDATE kanji SET kanji = COALESCE($1, kanji), on_reading = COALESCE($2, on_reading), kun_reading = COALESCE($3, kun_reading), kanji_rating = COALESCE($4, kanji_rating) WHERE id = $5 RETURNING id"
	_, err = tx.ExecContext(ctx, updateKanjiStatement, k.Kanji, k.OnReading, k.KunReading, k.KanjiRating, k.ID)
	if err != nil {
		return nil, err
	}

	// update kanji_sentence
	updateKanjiSentenceStatement := "UPDATE kanji_sentence SET example_sentence = COALESCE($1, example_sentence) WHERE sentence_id = $2"
	for _, exampleSentence := range k.ExampleSentences {
		_, err = tx.ExecContext(ctx, updateKanjiSentenceStatement, exampleSentence.ExampleSentence, exampleSentence.SentenceID)
		if err != nil {
			return nil, err
		}
	}

	//update kanji_example
	updateKanjiExampleStatement := "UPDATE kanji_example SET example_word = COALESCE($1, example_word) WHERE word_id = $2"
	for _, exampleWord := range k.ExampleWords {
		_, err = tx.ExecContext(ctx, updateKanjiExampleStatement, exampleWord.ExampleWord, exampleWord.WordId)
		if err != nil {
			return nil, err
		}
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
