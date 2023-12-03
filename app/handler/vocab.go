package handler

import (
	"context"
	"mon-backend/domain/vocabulary"
)

type VocabHandler struct {
	repo vocabulary.VocabRepository
}

type NewVocab struct {
	Vocab            string
	Definitions      string
	ExampleSentences string
	PartsOfSpeech    string
	Kanji            string
	VocabRating      int
	Username         string
}

func NewVocabHandler(repo vocabulary.VocabRepository) VocabHandler {
	if repo == nil {
		panic("nil vocab repo")
	}

	return VocabHandler{repo}
}

func (h VocabHandler) AddVocab(ctx context.Context, newVocab NewVocab) (*vocabulary.Vocab, error) {

	v, err := vocabulary.NewVocab(newVocab.Vocab, newVocab.Definitions, newVocab.Definitions, newVocab.PartsOfSpeech, newVocab.Kanji, newVocab.VocabRating, newVocab.Username)

	if err != nil {
		return nil, err
	}

	v, err = h.repo.AddVocab(ctx, v)

	if err != nil {
		return nil, err
	}

	return v, nil
}

func (h VocabHandler) GetVocab(ctx context.Context, username string, limit int) ([]*vocabulary.Vocab, error) {
	vocab, err := h.repo.GetVocab(ctx, username, limit)

	if err != nil {
		return nil, err
	}

	return vocab, nil
}

func (h VocabHandler) UpdateVocab(ctx context.Context, vocab vocabulary.Vocab) (*vocabulary.Vocab, error) {
	updatedVocab, err := h.repo.UpdateVocab(ctx, vocab)

	if err != nil {
		return nil, err
	}

	return updatedVocab, nil
}

func (h VocabHandler) DeleteVocab(id string) error {
	err := h.repo.DeleteVocab(id)

	if err != nil {
		return err
	}

	return nil
}
