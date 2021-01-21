package consumer

import (
	"context"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	consumergrpc "gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/grpc/consumer"
	"google.golang.org/grpc"
)

type Servicer interface {
	server.GrpcServerController
	Wallet(consumerID string) (map[string]int, error)
}

type Service struct {
	server.GrpcServer
	consumergrpc.UnimplementedConsumerServer
}

func NewConsumerService() *Service {
	return &Service{}
}

func (consumerSvc *Service) Run(host string, port int) error {
	consumerSvc.Gs = grpc.NewServer()
	consumergrpc.RegisterConsumerServer(consumerSvc.Gs, consumerSvc)
	return consumerSvc.GrpcServer.Run(host, port)
}

func (consumerSvc *Service) Wallet(_ context.Context, request *consumergrpc.WalletRequest) (*consumergrpc.WalletResponse, error) {
	cdb, err := NewDb()
	if err != nil {
		return nil, err
	}
	defer cdb.Close()

	consumer, err := cdb.Read(request.ConsumerId)
	if err != nil {
		return nil, err
	}

	response := &consumergrpc.WalletResponse{
		Items: util.MapToInt32(consumer.Wallet),
	}

	return response, nil
}

func (consumerSvc *Service) AddToWallet(_ context.Context, request *consumergrpc.AddRequest) (*consumergrpc.EmptyResponse, error) {
	cdb, err := NewDb()
	if err != nil {
		return nil, err
	}
	defer cdb.Close()

	err = cdb.AddToWallet(request.ConsumerId, request.ItemId)
	if err != nil {
		return nil, err
	}

	return &consumergrpc.EmptyResponse{}, nil
}

func (consumerSvc *Service) RemoveFromWallet(_ context.Context, request *consumergrpc.RemoveRequest) (*consumergrpc.EmptyResponse, error) {
	cdb, err := NewDb()
	if err != nil {
		return nil, err
	}
	defer cdb.Close()

	err = cdb.RemoveFromWallet(request.ConsumerId, request.ItemId, int(request.Quantity))
	if err != nil {
		return nil, err
	}

	return &consumergrpc.EmptyResponse{}, nil
}
