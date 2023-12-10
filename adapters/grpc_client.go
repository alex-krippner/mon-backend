package adapters

import (
	"errors"
	monNlp "mon-backend/adapters/genproto/monNlpService"
	"os"

	"google.golang.org/grpc"
)

func NewMonNlpClient() (client monNlp.MonNlpServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("MON_NLP_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env MON_NLP_GRPC_ADDR")
	}

	conn, err := grpc.Dial(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return monNlp.NewMonNlpServiceClient(conn), conn.Close, nil
}
