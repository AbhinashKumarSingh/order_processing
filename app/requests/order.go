package requests

type OrderReq struct {
	OrderID     string   `json:"order_id"`
	UserID      string   `json:"user_id"`
	ItemsIDs    []string `json:"items_ids"`
	TotalAmount float64  `json:"total_amount"`
	Status      string   `json:"status"`
}
