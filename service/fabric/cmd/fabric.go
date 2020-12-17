package main

import (
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/pkg/di"
	"gitlab.com/adesso-turkey/loyalty-backend-microservices/service/fabric"
	"log"
)

func main() {
	conf := di.InitializeConfig()
	chaincodeService := fabric.NewChaincodeService()

	err := chaincodeService.Run("", conf.Services["fabric"].GrpcPort)
	if err != nil {
		log.Fatalf("Unable to serve with gRPC %v", err)
	}
}
