package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

const selectVocabById = `
SELECT vocabulary.id, vocabulary.vocab, vocabulary.kanji, vocabulary.vocab_rating, vocabulary.username, vd.vocabulary_definitions, vs.example_sentences, ps.parts_of_speech
FROM vocabulary
	LEFT JOIN (
		SELECT vd.vocab_id AS id, json_agg(
			json_build_object('definition', def, 'id', vd.id, 'vocabId', vocab_id)) AS vocabulary_definitions
		FROM vocabulary_definition vd
		GROUP BY vd.vocab_id
	) vd USING(id)
	LEFT JOIN (
		SELECT vs.vocab_id AS id, json_agg(
			json_build_object('exampleSentence', example_sentence, 'id', vs.id, 'vocabId', vocab_id)) AS example_sentences
		FROM vocabulary_sentence vs
		GROUP BY vs.vocab_id
	) vs USING (id)
	LEFT JOIN (
		SELECT ps.vocab_id AS id, json_agg(
			json_build_object('partOfSpeech', part_of_speech, 'vocabId', ps.vocab_id)) AS parts_of_speech
		FROM (
			select ps.part_of_speech, vp.vocab_id
			from parts_of_speech ps
			join vocab_part_of_speech vp
			on vp.part_of_speech_id = ps.id
		) AS ps
		GROUP BY ps.vocab_id
	) ps USING (id)
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

type VocabDefinitions []VocabDefinition
type VocabSentences []VocabSentence
type PartsOfSpeech []PartOfSpeech

type Vocab struct {
	ID               string           `json:"id,omitempty"`
	Vocab            string           `json:"vocab,omitempty"`
	Definitions      VocabDefinitions `json:"definitions,omitempty"`
	ExampleSentences VocabSentences   `json:"exampleSentences,omitempty"`
	PartsOfSpeech    PartsOfSpeech    `json:"partsOfSpeech,omitempty"`
	Kanji            string           `json:"kanji,omitempty"`
	VocabRating      int              `json:"vocabRating,omitempty"`
	Username         string           `json:"username,omitempty"`
}

type CreateVocabRequest struct {
	Vocab            string
	Definitions      VocabDefinitions
	ExampleSentences VocabSentences
	PartsOfSpeech    PartsOfSpeech
	Kanji            string
	VocabRating      int
	Username         string
}

type UpdateVocabRequest struct {
	ID               *string          `json:"id,omitempty"`
	Vocab            *string          `json:"vocab,omitempty"`
	Definitions      VocabDefinitions `json:"definitions,omitempty"`
	ExampleSentences VocabSentences   `json:"exampleSentences,omitempty"`
	Kanji            *string          `json:"kanji,omitempty"`
	VocabRating      *int             `json:"vocabRating,omitempty"`
}

func (d *VocabDefinitions) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &d)
}

func (vs *VocabSentences) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &vs)
}

func (ps *PartsOfSpeech) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &ps)
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

	insertVocabStatement := "INSERT INTO vocabulary(vocab, kanji, vocab_rating, username) VALUES($1,$2,$3,$4) RETURNING id;"
	insertedVocabRow := s.conn.QueryRowContext(ctx, insertVocabStatement, v.Vocab, v.Kanji, v.VocabRating, v.Username)
	var vocabId string
	err = insertedVocabRow.Scan(&vocabId)
	if err != nil {
		return nil, err
	}

	insertDefinitionStatement := "INSERT INTO vocabulary_definition(vocab_id, def) VALUES($1, $2);"
	for _, d := range v.Definitions {
		_, err = tx.ExecContext(ctx, insertDefinitionStatement, vocabId, d.Definition)
		if err != nil {
			return nil, err
		}
	}

	insertSentenceStatement := "INSERT INTO vocabulary_sentence(vocab_id, example_sentence) VALUES($1, $2);"
	for _, s := range v.ExampleSentences {
		_, err = tx.ExecContext(ctx, insertSentenceStatement, vocabId, s.ExampleSentence)
		if err != nil {
			return nil, err
		}
	}

	insertPartOfSpeechStatement := "INSERT INTO parts_of_speech(part_of_speech) VALUES($1) RETURNING id;"

	partOfSpeechIds := make([]string, 0)
	for _, p := range v.PartsOfSpeech {
		row := tx.QueryRowContext(ctx, insertPartOfSpeechStatement, p.PartOfSpeech)
		var id string
		if err := row.Scan(&id); err != nil {
			return nil, err
		}

		partOfSpeechIds = append(partOfSpeechIds, id)

		if err != nil {
			return nil, err
		}
	}

	if len(partOfSpeechIds) != 0 {
		insertVocabPartOfSpeechStatement := "INSERT INTO vocab_part_of_speech(vocab_id, part_of_speech_id) VALUES($1, $2);"
		for _, id := range partOfSpeechIds {
			_, err = tx.ExecContext(ctx, insertVocabPartOfSpeechStatement, vocabId, id)
			if err != nil {
				return nil, err
			}
		}

	}

	row := tx.QueryRowContext(ctx, selectVocabById, vocabId)
	vocab, err := ScanVocab(row)
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
	SELECT vocabulary.id, vocabulary.vocab, vocabulary.kanji, vocabulary.vocab_rating, vocabulary.username, vd.vocabulary_definitions, vs.example_sentences, ps.parts_of_speech
	FROM vocabulary
		LEFT JOIN (
			SELECT vd.vocab_id AS id, json_agg(
				json_build_object('definition', def, 'id', vd.id, 'vocabId', vocab_id)) AS vocabulary_definitions
			FROM vocabulary_definition vd
			GROUP BY vd.vocab_id
		) vd USING(id)
		LEFT JOIN (
			SELECT vs.vocab_id AS id, json_agg(
				json_build_object('exampleSentence', example_sentence, 'id', vs.id, 'vocabId', vocab_id)) AS example_sentences
			FROM vocabulary_sentence vs
			GROUP BY vs.vocab_id
		) vs USING (id)
		LEFT JOIN (
			SELECT ps.vocab_id AS id, json_agg(
				json_build_object('partOfSpeech', part_of_speech, 'vocabId', ps.vocab_id)) AS parts_of_speech
			FROM (
				select ps.part_of_speech, vp.vocab_id
				from parts_of_speech ps
				join vocab_part_of_speech vp
				on vp.part_of_speech_id = ps.id
			) AS ps
			GROUP BY ps.vocab_id
		) ps USING (id);`

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

	updateVocabStatement := "UPDATE vocabulary SET vocab = COALESCE($1, vocab), kanji = COALESCE($2, kanji), vocab_rating = COALESCE($3, vocab_rating) WHERE id = $4;"
	_, err = tx.ExecContext(ctx, updateVocabStatement, v.Vocab, v.Kanji, v.VocabRating, v.ID)
	if err != nil {
		return nil, err
	}

	updateVocabDefStatement := "UPDATE vocabulary_definition SET def = COALESCE($1, def) WHERE id = $2;"
	for _, d := range v.Definitions {
		_, err = tx.ExecContext(ctx, updateVocabDefStatement, d.Definition, d.ID)
		if err != nil {
			return nil, err
		}
	}

	updateVocabSentenceStatement := "UPDATE vocabulary_sentence SET example_sentence = COALESCE($1, example_sentence) WHERE id = $2;"
	for _, vs := range v.ExampleSentences {
		_, err = tx.ExecContext(ctx, updateVocabSentenceStatement, vs.ExampleSentence, vs.ID)
		if err != nil {
			return nil, err
		}
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
