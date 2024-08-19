package main

import (
	"log"

	"github.com/SmoothWay/microservices/order/config"
	"github.com/SmoothWay/microservices/order/internal/adapters/db"
	"github.com/SmoothWay/microservices/order/internal/adapters/grpc"
	"github.com/SmoothWay/microservices/order/internal/adapters/payment"
	"github.com/SmoothWay/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDSN())
	if err != nil {
		log.Fatalf("failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("failed to initialize payment stub. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())

	grpcAdapter.Run()
}
