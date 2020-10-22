package server

import (
	"fmt"
	"net"

	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
)

type GrpcServerController interface {
	Run(host string, port int) error
}

type GrpcServer struct {
	Gs *grpc.Server
}

func (srv *GrpcServer) Run(host string, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Unable to listen %v", err)
	}

	log.Print("Starting [Item] gRPC server...")

	// srv.gs = grpc.NewServer()

	return srv.Gs.Serve(lis)
}
