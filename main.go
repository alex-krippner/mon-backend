package main

import (
	"context"
	"net/http"

	"mon-backend/ports"
	"mon-backend/server"
	"mon-backend/service"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx := context.Background()
	app, cleanup := service.NewApplication(ctx)
	defer cleanup()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(ports.NewHttpServer(app), router)
	})
}
