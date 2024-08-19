package api

import (
	"context"

	"github.com/SmoothWay/microservices/payment/internal/application/core/domain"
	"github.com/SmoothWay/microservices/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	err := a.db.Save(context.Background(), &payment)
	if err != nil {
		return domain.Payment{}, err
	}

	return payment, nil
}
