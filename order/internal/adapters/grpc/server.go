package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/SmoothWay/microservices-proto/golang/order"
	"github.com/SmoothWay/microservices/order/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api  ports.APIPort
	port int
	order.UnimplementedOrderServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port %v", a.port)
	}
}
