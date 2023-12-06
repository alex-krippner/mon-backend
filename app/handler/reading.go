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
	Username    string
}

func NewReadingHandler(repo domain.ReadingRepository) ReadingHandler {
	if repo == nil {
		panic("nil reading repo")
	}

	return ReadingHandler{repo}
}

func (h ReadingHandler) TranslateReading(ctx context.Context, newReading NewReading) (*domain.Reading, error) {
	r, err := domain.NewReading(newReading.Translation, newReading.Japanese, newReading.Title, newReading.Username)
	if err != nil {
		return nil, err
	}

	r, err = h.repo.CreateReading(ctx, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (h ReadingHandler) GetAllReading(ctx context.Context, username string) ([]*domain.Reading, error) {
	readings, err := h.repo.GetAllReading(ctx, username)

	if err != nil {
		return nil, err
	}

	return readings, nil
}

func (h ReadingHandler) DeleteReading(ctx context.Context, id string) error {
	err := h.repo.DeleteReading(id)

	if err != nil {
		return err
	}

	return nil
}

func (h ReadingHandler) UpdateReading(ctx context.Context, reading domain.Reading) (*domain.Reading, error) {
	updatedReading, err := h.repo.UpdateReading(ctx, reading)

	if err != nil {
		return nil, err
	}

	return updatedReading, nil
}
