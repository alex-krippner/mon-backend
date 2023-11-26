package kanji

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

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

type KanjiRepository interface {
	AddKanji(ctx context.Context, req *Kanji) (*Kanji, error)
	GetAllKanji(ctx context.Context, username string) ([]*Kanji, error)
	GetKanji(ctx context.Context, id string) (*Kanji, error)
	UpdateKanji(ctx context.Context, kanji Kanji) (*Kanji, error)
	DeleteKanji(id string) error
}

func NewKanji(exampleSentences string, exampleWords string, kanji string, kanjiRating int, kunReading string, meanings string, onReading, username string) (*Kanji, error) {
	if kanji == "" {
		return nil, errors.New("kanji is missing")
	}

	return &Kanji{
		ID:               uuid.New().String(),
		ExampleSentences: exampleSentences,
		ExampleWords:     exampleWords,
		Kanji:            kanji,
		KanjiRating:      kanjiRating,
		KunReading:       kunReading,
		Meanings:         meanings,
		OnReading:        onReading,
		Username:         username,
	}, nil
}
