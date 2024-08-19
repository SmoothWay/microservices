package main

import (
	"log"

	"github.com/SmoothWay/microservices/payment/config"
	"github.com/SmoothWay/microservices/payment/internal/adapters/db"
	"github.com/SmoothWay/microservices/payment/internal/adapters/grpc"
	"github.com/SmoothWay/microservices/payment/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
