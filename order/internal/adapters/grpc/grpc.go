package grpc

import (
	"context"

	"github.com/SmoothWay/microservices-proto/golang/order"
	"github.com/SmoothWay/microservices/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	var orderItems []domain.OrderItem
	for _, orderitem := range request.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: orderitem.ProductCode,
			UnitPrice:   orderitem.UnitPrice,
			Quantity:    orderitem.Quantity,
		})
	}
	newOrder := domain.NewOrder(request.UserId, orderItems)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &order.CreateOrderResponse{
		OrderId: result.ID,
	}, nil
}
