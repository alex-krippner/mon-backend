package adapters

import (
	"context"
	"database/sql"
	"fmt"
	"mon-backend/domain/vocabulary"
)

const selectVocabById = `
SELECT vocabulary.id, vocabulary.vocab, vocabulary.kanji, vocabulary.vocab_rating, vocabulary.username, vocabulary.definitions, vocabulary.example_sentences, vocabulary.parts_of_speech
FROM vocabulary
WHERE vocabulary.id = $1
	;`

type VocabRepository struct {
	db *sql.DB
}

func NewVocabRepository(db *sql.DB) *VocabRepository {
	return &VocabRepository{
		db: db,
	}
}

func ScanVocab(s Scanner) (*vocabulary.Vocab, error) {
	v := &vocabulary.Vocab{}
	if err := s.Scan(&v.ID, &v.Vocab, &v.Kanji, &v.VocabRating, &v.Username, &v.Definitions, &v.ExampleSentences, &v.PartsOfSpeech); err != nil {
		return nil, err
	}

	return v, nil
}

func (r VocabRepository) AddVocab(ctx context.Context, v *vocabulary.Vocab) (*vocabulary.Vocab, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	insertVocabStatement := "INSERT INTO vocabulary(vocab, kanji, vocab_rating, username, definitions, example_sentences, parts_of_speech) VALUES($1,$2,$3,$4,$5,$6,$7) RETURNING id, vocab, kanji, vocab_rating, username, definitions, example_sentences, parts_of_speech;"
	insertedVocabRow := r.db.QueryRowContext(ctx, insertVocabStatement, v.Vocab, v.Kanji, v.VocabRating, v.Username, v.Definitions, v.ExampleSentences, v.PartsOfSpeech)

	vocab, err := ScanVocab(insertedVocabRow)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return vocab, nil
}

func (r VocabRepository) GetVocab(ctx context.Context, username string, limit int) ([]*vocabulary.Vocab, error) {

	const selectStatement = `
	SELECT vocabulary.id, vocabulary.vocab, vocabulary.kanji, vocabulary.vocab_rating, vocabulary.username, vocabulary.definitions, vocabulary.example_sentences, vocabulary.parts_of_speech
	FROM vocabulary
	WHERE vocabulary.username = $1
	LIMIT $2
	;`

	rows, err := r.db.QueryContext(ctx, selectStatement, username, limit)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve vocab: %w", err)
	}

	defer rows.Close()

	var vocabSlice []*vocabulary.Vocab
	for rows.Next() {
		vocab, err := ScanVocab(rows)
		if err != nil {
			return nil, fmt.Errorf("could not scan vocab: %w", err)
		}

		vocabSlice = append(vocabSlice, vocab)
	}

	return vocabSlice, nil
}

func (r VocabRepository) UpdateVocab(ctx context.Context, vocab vocabulary.Vocab) (*vocabulary.Vocab, error) {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	updateVocabStatement := "UPDATE vocabulary SET vocab = COALESCE($1, vocab), kanji = COALESCE($2, kanji), vocab_rating = COALESCE($3, vocab_rating), definitions = COALESCE($4, definitions), example_sentences = COALESCE($5, example_sentences), parts_of_speech = COALESCE($6, parts_of_speech) WHERE id = $7;"
	_, err = tx.ExecContext(ctx, updateVocabStatement, vocab.Vocab, vocab.Kanji, vocab.VocabRating, vocab.Definitions, vocab.ExampleSentences, vocab.PartsOfSpeech, vocab.ID)
	if err != nil {
		return nil, err
	}

	row := tx.QueryRowContext(ctx, selectVocabById, vocab.ID)
	updatedVocab, err := ScanVocab(row)
	if err != nil {
		return nil, fmt.Errorf("could not scan vocab: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return updatedVocab, nil
}

func (r VocabRepository) DeleteVocab(id string) error {

	_, err := r.db.Exec("DELETE FROM vocabulary WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
