package apiserver

import (
	"context"
	"errors"
	"fmt"
	"mon-backend/storage"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var defaultStopTimeout = time.Second * 30

type APIServer struct {
	addr string
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
		addr: addr,
		storage: storage,
	}, nil
}

func (s *APIServer) Start(stop <-chan struct{}) error {
	srv := &http.Server{
		Addr: s.addr,
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
 defer cancel()

 return srv.Shutdown(ctx)
}

func (s *APIServer) router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", s.defaultRoute)
	router.Methods("POST").Path("/items").Handler(Endpoint{s.createItem})
	router.Methods("GET").Path("/items").Handler(Endpoint{s.listItems})
	return router
}

func (s *APIServer) defaultRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}