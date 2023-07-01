package vocabulary

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

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

type VocabRepository interface {
	AddVocab(ctx context.Context, req *Vocab) (*Vocab, error)
	GetVocab(ctx context.Context, username string, limit int) ([]*Vocab, error)
	UpdateVocab(ctx context.Context, vocab Vocab) (*Vocab, error)
	DeleteVocab(id string) error
}

func NewVocab(vocab string, definitions string, exampleSentences string, partsOfSpeech string, kanji string, vocabRating int, username string) (*Vocab, error) {
	if vocab == "" {
		return nil, errors.New("vocabulary is missing")
	}
	if username == "" {
		return nil, errors.New("username is missing")
	}

	return &Vocab{
		ID:               uuid.New().String(),
		Vocab:            vocab,
		Definitions:      definitions,
		ExampleSentences: exampleSentences,
		PartsOfSpeech:    partsOfSpeech,
		Kanji:            kanji,
		VocabRating:      vocabRating,
		Username:         username,
	}, nil
}
