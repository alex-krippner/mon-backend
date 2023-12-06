package app

import (
	"mon-backend/app/handler"
)

type Application struct {
	Handlers Handlers
}

type Handlers struct {
	ReadingHandler handler.ReadingHandler
	KanjiHandler   handler.KanjiHandler
}
