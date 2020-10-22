package client

import (
	"context"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/util"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/grpc/consumer"
	"google.golang.org/grpc"
)

type ConsumerServicer interface {
	Connect() error
	Close() error
	Wallet(consumerID string) (map[string]int, error)
	AddToWallet(consumerID, itemID string) error
	RemoveFromWallet(consumerID, itemID string, quantity int) error
}

type ConsumerService struct {
	conn   *grpc.ClientConn
	client consumer.ConsumerClient
}

func NewConsumerService() *ConsumerService {
	return &ConsumerService{}
}

func (consumerSvc *ConsumerService) Connect() error {
	conn, err := grpc.Dial("consumer:9012", grpc.WithInsecure())
	if err != nil {
		return err
	}

	consumerSvc.conn = conn
	consumerSvc.client = consumer.NewConsumerClient(consumerSvc.conn)

	return nil
}

func (consumerSvc *ConsumerService) Close() error {
	return consumerSvc.conn.Close()
}

func (consumerSvc *ConsumerService) Wallet(consumerID string) (map[string]int, error) {
	request := &consumer.WalletRequest{ConsumerId: consumerID}

	response, err := consumerSvc.client.Wallet(context.Background(), request)
	if err != nil {
		return nil, err
	}

	return util.MapToInt(response.Items), nil
}

func (consumerSvc *ConsumerService) AddToWallet(consumerID, itemID string) error {
	request := &consumer.AddRequest{ConsumerId: consumerID, ItemId: itemID}
	_, err := consumerSvc.client.AddToWallet(context.Background(), request)
	return err
}

func (consumerSvc *ConsumerService) RemoveFromWallet(consumerID, itemID string, quantity int) error {
	request := &consumer.RemoveRequest{ConsumerId: consumerID, ItemId: itemID, Quantity: int32(quantity)}
	_, err := consumerSvc.client.RemoveFromWallet(context.Background(), request)
	return err
}
