package storage

import (
	"context"
	"fmt"
)

type Vocab struct {
	ID          string
	Vocab       string
	Definition  string
	Kanji       string
	VocabRating int
	Username    string
}

type CreateVocabRequest struct {
	Vocab       string
	Definition  string
	Kanji       string
	VocabRating int
	Username    string
}

type UpdateVocabRequest struct {
	ID          *string `json:"id,omitempty"`
	Vocab       *string `json:"vocab,omitempty"`
	Definition  *string `json:"definition,omitempty"`
	Kanji       *string `json:"kanji,omitempty"`
	VocabRating *int    `json:"vocabRating,omitempty"`
}

func ScanVocab(s Scanner) (*Vocab, error) {
	v := &Vocab{}
	if err := s.Scan(&v.ID, &v.Vocab, &v.Definition, &v.Kanji, &v.VocabRating, &v.Username); err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Storage) CreateVocab(ctx context.Context, v CreateVocabRequest) (*Vocab, error) {
	insertStatement := "INSERT INTO vocabulary(vocab, def, kanji, vocab_rating, username) VALUES($1,$2,$3,$4,$5) RETURNING id, vocab, def, kanji, vocab_rating, username"
	row := s.conn.QueryRowContext(ctx, insertStatement, v.Vocab, v.Definition, v.Kanji, v.VocabRating, v.Username)
	return ScanVocab(row)
}

func (s *Storage) GetVocab(ctx context.Context, id string) (*Vocab, error) {

	selectStatement := "SELECT * FROM vocabulary WHERE id = $1"
	row := s.conn.QueryRowContext(ctx, selectStatement, id)
	return ScanVocab(row)
}

func (s *Storage) GetAllVocab(ctx context.Context) ([]*Vocab, error) {
	selectStatement := "SELECT * FROM vocabulary"
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
	updateStatement := "UPDATE vocabulary SET vocab = COALESCE($1, vocab), def = COALESCE($2, def), kanji = COALESCE($3, kanji), vocab_rating = COALESCE($4, vocab_rating) WHERE id = $5 RETURNING *"
	row := s.conn.QueryRowContext(ctx, updateStatement, v.Vocab, v.Definition, v.Kanji, v.VocabRating, v.ID)

	vocab, err := ScanVocab(row)

	if err != nil {
		return nil, fmt.Errorf("could not scan vocab: %w", err)
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
