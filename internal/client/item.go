package client

import (
	"context"
	"fmt"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"

	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/grpc/item"
	"google.golang.org/grpc"
)

type ItemServicer interface {
	Connect() error
	Close() error
	Create(model *item.CreateItemDto) error
	ItemExists(id string) (bool, error)
}

type ItemService struct {
	conn   *grpc.ClientConn
	client item.ItemClient
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (itemSvc *ItemService) Connect() error {
	conf := di.InitializeConfig()
	target := fmt.Sprintf("%s:%d", conf.Services["item"].Host, conf.Services["item"].GrpcPort)

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return err
	}

	itemSvc.conn = conn
	itemSvc.client = item.NewItemClient(itemSvc.conn)

	return nil
}

func (itemSvc *ItemService) Close() error {
	return itemSvc.conn.Close()
}

func (itemSvc *ItemService) Create(model *item.CreateItemDto) error {
	_, err := itemSvc.client.Create(context.Background(), model)

	return err
}

func (itemSvc *ItemService) ItemExists(id string) (bool, error) {
	response, err := itemSvc.client.ItemExists(context.Background(), &item.ItemExistsRequest{ID: id})
	if err != nil {
		return false, err
	}

	return response.Found, nil
}
