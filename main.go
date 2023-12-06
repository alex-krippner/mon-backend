package main

import (
	"context"
	"net/http"

	"mon-backend/ports"
	"mon-backend/server"
	"mon-backend/service"

	"github.com/go-chi/chi/v5"
)

const (
	apiServerAddrFlagName       string = "API_SERVER_ADDR"
	apiServerStorageDatabaseURL string = "DATABASE_URL"
)

func main() {
	ctx := context.Background()
	app := service.NewApplication(ctx)

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}
