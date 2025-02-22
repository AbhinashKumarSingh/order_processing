package models

type OrderItems struct {
	OrderID string `json:"order_id"`
	ItemId  string `json:"item_id"`
}

func (OrderItems) TableName() string {
	return "order_items"
}
