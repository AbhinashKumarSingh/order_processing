package models

import "time"

type Orders struct {
	OrderID     string     `json:"order_id"`
	UserID      string     `json:"user_id"`
	TotalAmount float64    `json:"total_amount"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

func (Orders) TableName() string {
	return "orders"
}
