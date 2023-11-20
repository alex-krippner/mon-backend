package reading

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Reading struct {
	ID          string `json:"id,omitempty"`
	Translation string `json:"translation,omitempty"`
	Japanese    string `json:"japanese,omitempty"`
	Title       string `json:"title,omitempty"`
	Username    string `json:"username,omitempty"`
}

type ReadingRepository interface {
	CreateReading(ctx context.Context, req *Reading) (*Reading, error)
	GetAllReading(ctx context.Context, username string) ([]*Reading, error)
	UpdateReading(ctx context.Context, req Reading) (*Reading, error)
	DeleteReading(id string) error
}

func NewReading(translation string, japanese string, title string, username string) (*Reading, error) {
	if translation == "" {
		return nil, errors.New("reading translation missing")
	}
	if japanese == "" {
		return nil, errors.New("japanese text missing")
	}
	if title == "" {
		return nil, errors.New("reading title is missing")
	}
	if username == "" {
		return nil, errors.New("username is missing")
	}

	return &Reading{
		ID:          uuid.New().String(),
		Translation: translation,
		Japanese:    japanese,
		Title:       title,
		Username:    username,
	}, nil
}