package handler

import (
	"context"
	"mon-backend/domain/reading"
)

type TokenizeResponse struct {
	Tokens []reading.Token
}

type MonNlpService interface {
	TokenizeText(ctx context.Context, text string) (TokenizeResponse, error)
}
