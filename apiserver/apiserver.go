package apiserver

import (
	"context"
	"errors"
	"fmt"
	"mon-backend/storage"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var defaultStopTimeout = time.Second * 30

type APIServer struct {
	addr    string
	storage *storage.Storage
}

type Endpoint struct {
	handler EndpointFunc
}

type EndpointFunc func(w http.ResponseWriter, req *http.Request) error

func (e Endpoint) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if err := e.handler(w, req); err != nil {
		println("could not process request", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error"))
	}
}

func NewAPIServer(addr string, storage *storage.Storage) (*APIServer, error) {
	if addr == "" {
		return nil, errors.New("addr cannot be blank")
	}

	return &APIServer{
		addr:    addr,
		storage: storage,
	}, nil
}

func (s *APIServer) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr:    s.addr,
		Handler: s.router(),
	}

	go func() {
		fmt.Println("Starting server for " + srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), defaultStopTimeout)
	/*
	 * Noob comment: The context needs to be cancelled to prevent a memory leak.
	 * Failing to cancel a context leads to the goroutine created by WithTimeout to be retained indefinitely
	 * https://go.dev/src/context/context.go?s=9162:9288
	 */
	defer cancel()

	return srv.Shutdown(ctx)
}

func (s *APIServer) router() http.Handler {
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	router.HandleFunc("/", s.defaultRoute)
	s.addKanjiHandlers(router)
	s.addVocabHandlers(router)
	s.addReadingHandlers(router)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8081"})

	corsHandler := handlers.CORS(headers, methods, origins)(router)
	return corsHandler
}

func (s *APIServer) defaultRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
