package db

import (
	"context"
	"fmt"

	"github.com/SmoothWay/microservices/payment/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Payment, error) {
	var paymentEntity Payment
	res := a.db.WithContext(ctx).First(&paymentEntity, id)
	payment := domain.Payment{
		ID:         int64(paymentEntity.ID),
		CUstomerID: paymentEntity.CustomerID,
		Status:     paymentEntity.Status,
		OrderId:    paymentEntity.OrderID,
		TotalPrice: paymentEntity.TotalPrice,
		CreatedAt:  paymentEntity.CreatedAt.UnixNano(),
	}

	return payment, res.Error
}

func (a Adapter) Save(ctx context.Context, payment *domain.Payment) error {
	orderModel := Payment{
		CustomerID: payment.CUstomerID,
		Status:     payment.Status,
		OrderID:    payment.OrderId,
		TotalPrice: payment.TotalPrice,
	}

	res := a.db.WithContext(ctx).Create(&orderModel)
	if res.Error == nil {
		payment.ID = int64(orderModel.ID)
	}

	return res.Error
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	err := db.AutoMigrate(&Payment{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}
