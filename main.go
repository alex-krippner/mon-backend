package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"mon-backend/apiserver"
	"mon-backend/storage"
)

const (
	apiServerAddrFlagName       string = "API_SERVER_ADDR"
	apiServerStorageDatabaseURL string = "DATABASE_URL"
)

func main() {
	if err := startServer(); err != nil {
		fmt.Println("could not run application", err)
	}
}

func startServer() error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// Noob comment: empty struct requires no memory
	stopper := make(chan struct{})
	go func() {
		// Noob Comment: Block here until a value is received by the channel "done"
		<-done
		close(stopper)
	}()
	databaseURL := os.Getenv(apiServerStorageDatabaseURL)
	s, err := storage.NewStorage(databaseURL)
	if err != nil {
		return fmt.Errorf("could not initialize storage: %w", err)
	}

	addr := os.Getenv(apiServerAddrFlagName)
	server, err := apiserver.NewAPIServer(addr, s)
	if err != nil {
		return err
	}

	return server.Start(stopper)
}
