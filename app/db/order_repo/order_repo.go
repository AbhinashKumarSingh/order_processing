package order_repo

import (
	"example.com/m/v2/config"
	"example.com/m/v2/models"
	"example.com/m/v2/response"
)

func CreateOrder(order *models.Orders) error {

	return config.DB.
		Model(&models.Orders{}).
		Create(order).Error

}

func GetOrderStatus(orderID string) (string, error) {

	var status string
	err := config.DB.
		Model(&models.Orders{}).
		Select("status").
		Where("order_id = ?", orderID).
		First(&status).Error

	return status, err
}

func UpdateStatus(orderID string, updates map[string]interface{}) error {
	return config.DB.
		Model(&models.Orders{}).
		Where("order_id = ?", orderID).
		Updates(updates).Error

}

func TotalOrderProcessed() (int64, error) {
	var count int64
	err := config.DB.
		Model(&models.Orders{}).
		Count(&count).Error
	return count, err
}

func AvgOrderProcessingTime() (float64, error) {
	var avgProcessingTime float64
	err := config.DB.
		Raw(`SELECT 
				AVG(TIMESTAMPDIFF(SECOND, created_at, completed_at)) 
			 FROM orders 
			 WHERE status = 'Completed';`).
		Scan(&avgProcessingTime).Error

	return avgProcessingTime, err
}

func OrderStatusByCount() ([]*response.OrderStatusByCount, error) {

	ordersStatusByCount := make([]*response.OrderStatusByCount, 0)
	err := config.DB.
		Raw(`SELECT 
					status, 
					COUNT(*) as count 
				FROM orders 
				GROUP BY status`).
		Find(&ordersStatusByCount).Error

	return ordersStatusByCount, err
}
