package domain

import "time"

type Payment struct {
	ID         int64   `json:"id"`
	CUstomerID int64   `json:"customer_id"`
	Status     string  `json:"status"`
	OrderId    int64   `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
	CreatedAt  int64   `json:"created_at"`
}

func NewPayment(customerId, orderId int64, totalPrice float32) Payment {
	return Payment{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CUstomerID: customerId,
		OrderId:    orderId,
		TotalPrice: totalPrice,
	}
}
