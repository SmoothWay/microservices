package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/SmoothWay/microservices-proto/golang/payment"
	"github.com/SmoothWay/microservices/payment/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api    ports.ApiPort
	port   int
	server *grpc.Server
	payment.UnimplementedPaymentServer
}

func NewAdapter(api ports.ApiPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	a.server = grpcServer
	payment.RegisterPaymentServer(grpcServer, a)

	log.Printf("starting payment service on port %d ...", a.port)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc server. Error: %v", err)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
