package order_items_repo

import (
	"example.com/m/v2/config"
	"example.com/m/v2/models"
)

func CreateOrderItems(orderItems []*models.OrderItems) error {

	return config.DB.
		Model(&models.OrderItems{}).
		Create(orderItems).Error

}
