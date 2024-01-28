package adapters

import (
	"context"
	monNlp "mon-backend/adapters/genproto/monNlpService"
	"mon-backend/app/handler"
	"mon-backend/domain/reading"

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

	tokens := mapTokens(res.Tokens)

	return handler.TokenizeResponse{Tokens: tokens}, nil
}

func mapTokens(responseTokens []*monNlp.Token) []reading.Token {
	var tokens []reading.Token

	for i := 0; i < len(responseTokens); i++ {
		t := reading.Token{
			Text:         responseTokens[i].Text,
			Lemma:        responseTokens[i].Lemma_,
			Normalized:   responseTokens[i].Norm_,
			Prefix:       responseTokens[i].Prefix_,
			Suffix:       responseTokens[i].Suffix_,
			PartOfSpeech: responseTokens[i].Pos_,
			Tag:          responseTokens[i].Tag_,
			Dependency:   responseTokens[i].Dep_,
			Lang:         responseTokens[i].Lang_,
		}

		tokens = append(tokens, t)
	}

	return tokens
}
