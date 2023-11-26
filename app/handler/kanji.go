package handler

import (
	"context"
	"mon-backend/domain/kanji"
)

type KanjiHandler struct {
	repo kanji.KanjiRepository
}

type NewKanji struct {
	ExampleSentences string
	ExampleWords     string
	Kanji            string
	KanjiRating      int
	KunReading       string
	Meanings         string
	OnReading        string
	Username         string
}

func NewKanjiHandler(repo kanji.KanjiRepository) KanjiHandler {
	if repo == nil {
		panic("nil kanji repo")
	}

	return KanjiHandler{repo}
}

func (h KanjiHandler) AddKanji(ctx context.Context, newKanji NewKanji) (*kanji.Kanji, error) {
	k, err := kanji.NewKanji(newKanji.ExampleSentences, newKanji.ExampleWords, newKanji.Kanji, newKanji.KanjiRating, newKanji.KunReading, newKanji.Meanings, newKanji.OnReading, newKanji.Username)

	if err != nil {
		return nil, err
	}

	k, err = h.repo.AddKanji(ctx, k)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func (h KanjiHandler) GetAllKanji(ctx context.Context, username string) ([]*kanji.Kanji, error) {
	kanjis, err := h.repo.GetAllKanji(ctx, username)

	if err != nil {
		return nil, err
	}

	return kanjis, nil
}

func (h KanjiHandler) GetKanji(ctx context.Context, id string) (*kanji.Kanji, error) {
	kanji, err := h.repo.GetKanji(ctx, id)

	if err != nil {
		return nil, err
	}

	return kanji, nil
}

func (h KanjiHandler) UpdateKanji(ctx context.Context, kanji kanji.Kanji) (*kanji.Kanji, error) {
	updatedKanji, err := h.repo.UpdateKanji(ctx, kanji)

	if err != nil {
		return nil, err
	}

	return updatedKanji, nil
}

func (h KanjiHandler) DeleteKanji(id string) error {
	err := h.repo.DeleteKanji(id)

	if err != nil {
		return err
	}

	return nil
}
