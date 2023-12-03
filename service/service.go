package service

import (
	"context"
	"mon-backend/adapters"
	"mon-backend/app"
	"mon-backend/app/handler"
)

func NewApplication(ctx context.Context) app.Application {
	return newApplication(ctx)
}

func newApplication(ctx context.Context) app.Application {
	db, err := adapters.GetDatabase()
	if err != nil {
		panic(err)
	}
	repositories := adapters.InitRepositories(db)

	return app.Application{
		Handlers: app.Handlers{
			ReadingHandler: handler.NewReadingHandler(repositories.ReadingRepository),
			KanjiHandler:   handler.NewKanjiHandler(repositories.KanjiRepository),
			VocabHandler:   handler.NewVocabHandler(repositories.VocabRepository),
		},
	}
}
