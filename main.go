package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"mon-backend/apiserver"
	"mon-backend/storage"

	"github.com/urfave/cli/v2"
)

const (
	apiServerAddrFlagName       string = "addr"
	apiServerStorageDatabaseURL string = "database-url"
)

func main() {
	if err := app().Run(os.Args); err != nil {
		fmt.Println("could not run application", err)
	}
}

func app() *cli.App {
	return &cli.App{
		Name:  "mon-api-server",
		Usage: "The API",
		Commands: []*cli.Command{
			apiServerCmd(),
		},
	}
}

func apiServerCmd() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "starts the API server",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: apiServerAddrFlagName, EnvVars: []string{"API_SERVER_ADDR"}},
			&cli.StringFlag{Name: apiServerStorageDatabaseURL, EnvVars: []string{"DATABASE_URL"}},
		},
		Action: func(c *cli.Context) error {
			done := make(chan os.Signal, 1)
			signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			// Noob comment: empty struct requires no memory
			stopper := make(chan struct{})
			go func() {
				// Noob Comment: Block here until a value is received by the channel "done"
				<-done
				close(stopper)
			}()

			databaseURL := c.String(apiServerStorageDatabaseURL)
			s, err := storage.NewStorage(databaseURL)
			if err != nil {
				return fmt.Errorf("could not initialize storage: %w", err)
			}

			addr := c.String(apiServerAddrFlagName)
			server, err := apiserver.NewAPIServer(addr, s)
			if err != nil {
				return err
			}

			return server.Start(stopper)
		},
	}
}
