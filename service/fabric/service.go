package fabric

import (
	"context"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/internal/server"
	ccgrpc "gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/grpc/chaincode"
	"google.golang.org/grpc"
)

type Servicer interface {
	server.GrpcServerController
}

type Service struct {
	server.GrpcServer
	ccgrpc.UnimplementedFabricChaincodeServer
}

func NewChaincodeService() *Service {
	return &Service{}
}

func (ccSvc *Service) Run(host string, port int) error {
	ccSvc.Gs = grpc.NewServer()
	ccgrpc.RegisterFabricChaincodeServer(ccSvc.Gs, ccSvc)
	return ccSvc.GrpcServer.Run(host, port)
}

func (ccSvc *Service) Up(_ context.Context, request *ccgrpc.ChaincodeRequest) (*ccgrpc.EmptyResponse, error) {
	/*	ccPack, err := CreateExternalConfig(request.Name)
		if err != nil {
			return nil, err
		}
	*/
	return &ccgrpc.EmptyResponse{}, nil
}
