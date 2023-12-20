package handler

import (
	"context"
)

type TokenizeResponse struct {
	Value []string
}

type MonNlpService interface {
	TokenizeText(ctx context.Context, text string) (TokenizeResponse, error)
}
