package ports

import (
	"context"

	"github.com/SmoothWay/microservices/payment/internal/application/core/domain"
)

type ApiPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
