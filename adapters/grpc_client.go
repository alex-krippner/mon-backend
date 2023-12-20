package adapters

import (
	"errors"
	monNlp "mon-backend/adapters/genproto/monNlpService"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMonNlpClient() (client monNlp.MonNlpServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("MON_NLP_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env MON_NLP_GRPC_ADDR")
	}
	logrus.Info("Creating GRPC client connection to the target " + grpcAddr)
	// TODO: Check if having credentials is necessary
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return monNlp.NewMonNlpServiceClient(conn), conn.Close, nil
}
