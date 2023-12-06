package handler

import (
	"context"
	"mon-backend/domain"
)

type ReadingHandler struct {
	repo domain.ReadingRepository
}

type NewReading struct {
	Translation string
	Japanese    string
	Title       string
}

func NewReadingHandler(repo domain.ReadingRepository) ReadingHandler {
	if repo == nil {
		panic("nil reading repo")
	}

	return ReadingHandler{repo}
}

func (h ReadingHandler) TranslateReading(ctx context.Context, newReading NewReading) (*domain.Reading, error) {
	r, err := domain.NewReading(newReading.Translation, newReading.Japanese, newReading.Title)
	if err != nil {
		return nil, err
	}

	r, err = h.repo.CreateReading(ctx, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (h ReadingHandler) DeleteReading(ctx context.Context, id string) error {
	err := h.repo.DeleteReading(id)

	if err != nil {
		return err
	}

	return nil
}
