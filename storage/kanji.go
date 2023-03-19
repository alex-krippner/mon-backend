package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

const selectKanijById = `
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username,
		(SELECT  json_build_object('id', kanji.id, 'exampleSentence', ks.example_sentence, 'kanjiId', ks.kanji_id)
			FROM kanji_sentence ks
			WHERE ks.kanji_id = kanji.id
		) AS example_sentences,
		ke.example_words, m.meanings
	FROM kanji
		LEFT JOIN (
			SELECT ke.kanji_id AS id, json_agg(
				json_build_object('exampleWord', example_word, 'id', ke.id, 'kanjiId', kanji_id)) AS example_words
			FROM kanji_example ke
			GROUP BY ke.kanji_id
		) ke USING (id)
		LEFT JOIN (
			SELECT m.kanji_id AS id, json_agg(
				json_build_object('meaning', meaning, 'kanjiId', m.kanji_id)) AS meanings
			FROM (
				SELECT m.meaning, km.kanji_id
				FROM meaning m
				JOIN kanji_meaning km
				ON km.meaning_id = m.id
			) AS m
			GROUP BY m.kanji_id
		) m USING (id)
	WHERE kanji.id = $1
		;`

type ExampleSentence struct {
	ID              string `json:"id,omitempty"`
	KanjiID         string `json:"kanjiId,omitempty"`
	ExampleSentence string `json:"exampleSentence,omitempty"`
}

type ExampleWord struct {
	ID          string `json:"id,omitempty"`
	KanjiID     string `json:"kanjiId,omitempty"`
	ExampleWord string `json:"exampleWord,omitempty"`
}

type Meaning struct {
	KanjiId string `json:"kanjiId,omitempty"`
	Meaning string `json:"meaning,omitempty"`
}

type ExampleWords []ExampleWord
type Meanings []Meaning

type Kanji struct {
	ID               string          `json:"id,omitempty"`
	ExampleSentences ExampleSentence `json:"exampleSentences,omitempty"`
	ExampleWords     ExampleWords    `json:"exampleWords,omitempty"`
	Kanji            string          `json:"kanji,omitempty"`
	KanjiRating      int             `json:"kanjiRating,omitempty"`
	KunReading       string          `json:"kunReading,omitempty"`
	Meanings         Meanings        `json:"meanings,omitempty"`
	OnReading        string          `json:"onReading,omitempty"`
	Username         string          `json:"username,omitempty"`
}

type CreateKanjiRequest struct {
	ExampleSentences ExampleSentence
	ExampleWords     ExampleWords
	Kanji            string
	KanjiRating      int
	KunReading       string
	Meanings         Meanings
	OnReading        string
	Username         string
}

type UpdateKanjiRequest struct {
	ID               *string         `json:"id,omitempty"`
	Kanji            *string         `json:"kanji,omitempty"`
	ExampleSentences ExampleSentence `json:"exampleSentences,omitempty"`
	ExampleWords     ExampleWords    `json:"exampleWords,omitempty"`
	OnReading        *string         `json:"onReading,omitempty"`
	KunReading       *string         `json:"kunReading,omitempty"`
	KanjiRating      *int            `json:"kanjiRating,omitempty"`
}

//TODO: Remove function duplication. Try composition, embedding, or factory
func (e *ExampleSentence) Scan(src interface{}) error {
	if src == nil {
		return HandleNil(e)
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")

	}
	return json.Unmarshal(b, &e)
}

func (e *ExampleWords) Scan(src interface{}) error {
	if src == nil {
		return HandleNil(e)
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")

	}
	return json.Unmarshal(b, &e)
}

func (m *Meanings) Scan(src interface{}) error {
	if src == nil {
		return HandleNil(m)
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &m)
}

func ScanKanji(s Scanner) (*Kanji, error) {
	k := &Kanji{}
	if err := s.Scan(&k.ID, &k.Kanji, &k.OnReading, &k.KunReading, &k.KanjiRating, &k.Username, &k.ExampleSentences, &k.ExampleWords, &k.Meanings); err != nil {
		return nil, err
	}

	return k, nil
}

func (s *Storage) CreateKanji(ctx context.Context, k CreateKanjiRequest) (*Kanji, error) {
	fmt.Println(k.Meanings)
	fmt.Println(k)
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
	_, err = tx.ExecContext(ctx, insertKanjiSentenceStatement, kanjiId, k.ExampleSentences.ExampleSentence)
	if err != nil {
		return nil, err
	}

	// insert into kanji_example
	insertKanjiExampleStatement := "INSERT INTO kanji_example(kanji_id, example_word) VALUES($1, $2);"
	for _, w := range k.ExampleWords {
		_, err = tx.ExecContext(ctx, insertKanjiExampleStatement, kanjiId, w.ExampleWord)
		if err != nil {
			return nil, err
		}
	}

	insertMeaningStatement := "INSERT INTO meaning(meaning) VALUES($1) RETURNING id;"

	meaningIds := make([]string, 0)
	for _, m := range k.Meanings {
		row := tx.QueryRowContext(ctx, insertMeaningStatement, m.Meaning)
		var id string
		if err := row.Scan(&id); err != nil {
			return nil, err
		}

		meaningIds = append(meaningIds, id)

		if err != nil {
			return nil, err
		}
	}
	fmt.Println(meaningIds)
	if len(meaningIds) != 0 {
		insertKanjiMeanings := "INSERT INTO kanji_meaning(kanji_id, meaning_id) VALUES($1, $2);"
		for _, id := range meaningIds {
			_, err = tx.ExecContext(ctx, insertKanjiMeanings, kanjiId, id)
			if err != nil {
				return nil, err
			}
		}
	}

	// select from kanji, kanji_sentence, and kanji_example
	row := tx.QueryRowContext(ctx, selectKanijById, kanjiId)
	kanji, err := ScanKanji(row)
	if err != nil {
		return nil, err
	}
	// commit transaction
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
	SELECT kanji.id, kanji.kanji, kanji.on_reading, kanji.kun_reading, kanji.kanji_rating, kanji.username,
		(SELECT  json_build_object('id', kanji.id, 'exampleSentence', ks.example_sentence, 'kanjiId', ks.kanji_id)
			FROM kanji_sentence ks
			WHERE ks.kanji_id = kanji.id
		) AS example_sentences,
		ke.example_words, m.meanings
	FROM kanji
		LEFT JOIN (
			SELECT ke.kanji_id AS id, json_agg(
				json_build_object('exampleWord', example_word, 'id', ke.id, 'kanjiId', kanji_id)) AS example_words
			FROM kanji_example ke
			GROUP BY ke.kanji_id
		) ke USING (id)
		LEFT JOIN (
			SELECT m.kanji_id AS id, json_agg(
				json_build_object('meaning', meaning, 'kanjiId', m.kanji_id)) AS meanings
			FROM (
				SELECT m.meaning, km.kanji_id
				FROM meaning m
				JOIN kanji_meaning km
				ON km.meaning_id = m.id
			) AS m
			GROUP BY m.kanji_id
		) m USING (id);
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
	updateKanjiSentenceStatement := "UPDATE kanji_sentence SET example_sentence = COALESCE($1, example_sentence) WHERE id = $2"
	_, err = tx.ExecContext(ctx, updateKanjiSentenceStatement, k.ExampleSentences.ExampleSentence, k.ID)
	if err != nil {
		return nil, err
	}

	//update kanji_example
	updateKanjiExampleStatement := "UPDATE kanji_example SET example_word = COALESCE($1, example_word) WHERE id = $2"
	for _, exampleWord := range k.ExampleWords {
		_, err = tx.ExecContext(ctx, updateKanjiExampleStatement, exampleWord.ExampleWord, exampleWord.ID)
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
