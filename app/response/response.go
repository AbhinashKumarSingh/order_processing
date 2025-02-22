package response

type OrderMetrics struct {
	TotalOrders       int64                 `json:"total_orders"`
	AvgProcessingTime float64               `json:"avg_processing_time"` // In seconds
	OrdersByStatus    []*OrderStatusByCount `json:"orders_by_status"`
}

type OrderStatusByCount struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}
