package handler

import (
	"context"
	"mon-backend/domain/reading"
)

type ReadingHandler struct {
	repo          reading.ReadingRepository
	monNlpService MonNlpService
}

type NewReading struct {
	Translation string
	Japanese    string
	Title       string
	Username    string
}

func NewReadingHandler(repo reading.ReadingRepository, monNlpService MonNlpService) ReadingHandler {
	if repo == nil {
		panic("nil reading repo")
	}

	return ReadingHandler{repo, monNlpService}
}

func (h ReadingHandler) TranslateReading(ctx context.Context, newReading NewReading) (*reading.Reading, error) {

	r, err := reading.NewReading(newReading.Translation, newReading.Japanese, newReading.Title, newReading.Username)

	if err != nil {
		return nil, err
	}

	r, err = h.repo.CreateReading(ctx, r)
	if err != nil {
		return nil, err
	}

	tokenizedText, err := h.monNlpService.TokenizeText(ctx, newReading.Japanese)

	if err != nil {
		return nil, err
	}
	r.Tokens = tokenizedText.Tokens

	return r, nil
}

func (h ReadingHandler) GetAllReading(ctx context.Context, username string) ([]*reading.Reading, error) {
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

func (h ReadingHandler) UpdateReading(ctx context.Context, reading reading.Reading) (*reading.Reading, error) {
	updatedReading, err := h.repo.UpdateReading(ctx, reading)

	if err != nil {
		return nil, err
	}

	return updatedReading, nil
}
