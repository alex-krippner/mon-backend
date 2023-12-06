package storage

import (
	"context"
	"fmt"
)

const selectVocabById = `
SELECT vocabulary.id, vocabulary.vocab, vocabulary.kanji, vocabulary.vocab_rating, vocabulary.username, vocabulary.definitions, vocabulary.example_sentences, vocabulary.parts_of_speech
FROM vocabulary
WHERE vocabulary.id = $1
	;`

type VocabDefinition struct {
	ID         string `json:"id,omitempty"`
	VocabID    string `json:"vocabId,omitempty"`
	Definition string `json:"definition,omitempty"`
}

type VocabSentence struct {
	ID              string `json:"id,omitempty"`
	VocabID         string `json:"vocabId,omitempty"`
	ExampleSentence string `json:"exampleSentence,omitempty"`
}

type PartOfSpeech struct {
	VocabID      string `json:"vocabId,omitempty"`
	PartOfSpeech string `json:"partOfSpeech"`
}

type Vocab struct {
	ID               string `json:"id,omitempty"`
	Vocab            string `json:"vocab,omitempty"`
	Definitions      string `json:"definitions,omitempty"`
	ExampleSentences string `json:"exampleSentences,omitempty"`
	PartsOfSpeech    string `json:"partsOfSpeech,omitempty"`
	Kanji            string `json:"kanji,omitempty"`
	VocabRating      int    `json:"vocabRating,omitempty"`
	Username         string `json:"username,omitempty"`
}

type CreateVocabRequest struct {
	Vocab            string
	Definitions      string
	ExampleSentences string
	PartsOfSpeech    string
	Kanji            string
	VocabRating      int
	Username         string
}

type UpdateVocabRequest struct {
	ID               *string `json:"id,omitempty"`
	Vocab            *string `json:"vocab,omitempty"`
	Definitions      *string `json:"definitions,omitempty"`
	ExampleSentences *string `json:"exampleSentences,omitempty"`
	PartsOfSpeech    *string `json:"partsOfSpeech,omitempty"`
	Kanji            *string `json:"kanji,omitempty"`
	VocabRating      *int    `json:"vocabRating,omitempty"`
}

func ScanVocab(s Scanner) (*Vocab, error) {
	v := &Vocab{}
	if err := s.Scan(&v.ID, &v.Vocab, &v.Kanji, &v.VocabRating, &v.Username, &v.Definitions, &v.ExampleSentences, &v.PartsOfSpeech); err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Storage) CreateVocab(ctx context.Context, v CreateVocabRequest) (*Vocab, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	insertVocabStatement := "INSERT INTO vocabulary(vocab, kanji, vocab_rating, username, definitions, example_sentences, parts_of_speech) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, vocab, kanji, vocab_rating, username, definitions, example_sentences, parts_of_speech;"
	insertedVocabRow := s.conn.QueryRowContext(ctx, insertVocabStatement, v.Vocab, v.Kanji, v.VocabRating, v.Username, v.Definitions, v.ExampleSentences, v.PartsOfSpeech)

	vocab, err := ScanVocab(insertedVocabRow)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return vocab, nil
}

func (s *Storage) GetVocab(ctx context.Context, id string) (*Vocab, error) {
	row := s.conn.QueryRowContext(ctx, selectVocabById, id)
	return ScanVocab(row)
}

func (s *Storage) GetAllVocab(ctx context.Context) ([]*Vocab, error) {

	const selectStatement = `
	SELECT vocabulary.id, vocabulary.vocab, vocabulary.kanji, vocabulary.vocab_rating, vocabulary.username, vocabulary.definitions, vocabulary.example_sentences, vocabulary.parts_of_speech
	FROM vocabulary;`

	rows, err := s.conn.QueryContext(ctx, selectStatement)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve vocab: %w", err)
	}

	defer rows.Close()

	var vocabSlice []*Vocab
	for rows.Next() {
		vocab, err := ScanVocab(rows)
		if err != nil {
			return nil, fmt.Errorf("could not scan vocab: %w", err)
		}

		vocabSlice = append(vocabSlice, vocab)
	}

	return vocabSlice, nil
}

func (s *Storage) UpdateVocab(ctx context.Context, v UpdateVocabRequest) (*Vocab, error) {
	tx, err := s.conn.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	updateVocabStatement := "UPDATE vocabulary SET vocab = COALESCE($1, vocab), kanji = COALESCE($2, kanji), vocab_rating = COALESCE($3, vocab_rating), definitions = COALESCE($4, definitions), example_sentences = COALESCE($5, example_sentences), parts_of_speech = COALESCE($6, parts_of_speech) WHERE id = $7;"
	_, err = tx.ExecContext(ctx, updateVocabStatement, v.Vocab, v.Kanji, v.VocabRating, v.Definitions, v.ExampleSentences, v.PartsOfSpeech, v.ID)
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(ctx, selectVocabById, v.ID)
	vocab, err := ScanVocab(row)
	if err != nil {
		return nil, fmt.Errorf("could not scan vocab: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return vocab, nil
}

func (s *Storage) DeleteVocab(id string) error {
	_, err := s.conn.Exec("DELETE FROM vocabulary WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
