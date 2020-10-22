package item

import (
	"context"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"

	itemgrpc "gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/grpc/item"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type Servicer interface {
	server.GrpcServerController
	Create(model *CreateDto) error
}

type Service struct {
	server.GrpcServer
}

func NewItemService() *Service {
	return &Service{}
}

func (itemSvc *Service) Run(host string, port int) error {
	itemSvc.Gs = grpc.NewServer()
	itemgrpc.RegisterItemServer(itemSvc.Gs, itemSvc)
	return itemSvc.GrpcServer.Run(host, port)
}

func (itemSvc *Service) Create(_ context.Context, model *itemgrpc.CreateItemDto) (*itemgrpc.CreateResponse, error) {
	idb, err := NewDb()
	if err != nil {
		return nil, err
	}
	defer idb.Close()

	company, err := primitive.ObjectIDFromHex(model.Company)
	if err != nil {
		return nil, err
	}

	product, err := primitive.ObjectIDFromHex(model.Product)
	if err != nil {
		return nil, err
	}

	_, err = idb.Create(&CreateDto{
		Name:    model.Name,
		Company: company,
		Product: product,
		Point:   int(model.Point),
		Code:    model.Code,
	})
	if err != nil {
		return nil, err
	}

	return &itemgrpc.CreateResponse{
		Fail: false,
	}, nil
}

func (itemSvc *Service) ItemExists(_ context.Context, request *itemgrpc.ItemExistsRequest) (*itemgrpc.ItemExistsResponse, error) {
	idb, err := NewDb()
	if err != nil {
		return nil, err
	}

	entity, err := idb.Read(request.ID)
	if err != nil {
		return nil, err
	}

	return &itemgrpc.ItemExistsResponse{
		Found: entity != nil,
	}, nil
}
