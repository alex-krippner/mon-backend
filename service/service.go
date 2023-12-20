package service

import (
	"context"
	"mon-backend/adapters"
	"mon-backend/app"
	"mon-backend/app/handler"
)

func NewApplication(ctx context.Context) (app.Application, func()) {
	monNlpClient, closeMonNlpClient, err := adapters.NewMonNlpClient()

	if err != nil {
		panic(err)
	}

	monNlpGrpc := adapters.NewMonNlpGrpc(monNlpClient)

	return newApplication(ctx, monNlpGrpc),
		func() {
			_ = closeMonNlpClient()
		}
}

func newApplication(ctx context.Context, monNlpService handler.MonNlpService) app.Application {
	db, err := adapters.GetDatabase()
	if err != nil {
		panic(err)
	}
	repositories := adapters.InitRepositories(db)

	return app.Application{
		Handlers: app.Handlers{
			ReadingHandler: handler.NewReadingHandler(repositories.ReadingRepository, monNlpService),
			KanjiHandler:   handler.NewKanjiHandler(repositories.KanjiRepository),
			VocabHandler:   handler.NewVocabHandler(repositories.VocabRepository),
		},
	}
}
