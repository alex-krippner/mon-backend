package adapters

import (
	"context"
	monNlp "mon-backend/adapters/genproto/monNlpService"
	"mon-backend/app/handler"

	"github.com/pkg/errors"
)

type MonGrpc struct {
	client monNlp.MonNlpServiceClient
}

func NewMonNlpGrpc(client monNlp.MonNlpServiceClient) MonGrpc {
	return MonGrpc{client: client}
}

func (s MonGrpc) TokenizeText(ctx context.Context, text string) (handler.TokenizeResponse, error) {
	res, err := s.client.Tokenize(ctx, &monNlp.TokenizeRequest{Text: text})

	if err != nil {
		errors.Wrap(err, "unable to tokenize text")
	}

	return handler.TokenizeResponse{Value: res.GetTokens()}, nil
}
