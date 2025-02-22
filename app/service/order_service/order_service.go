package order_service

import (
	"example.com/m/v2/constants"
	"example.com/m/v2/db/order_items_repo"
	"example.com/m/v2/db/order_repo"
	"example.com/m/v2/models"
	"example.com/m/v2/requests"
	"example.com/m/v2/response"
	"github.com/phuslu/log"
)

func CreateOrderService(input requests.OrderReq) error {

	order := &Order{ID: input.OrderID, Status: constants.Pending}
	mu.Lock()
	orders[input.OrderID] = order
	mu.Unlock()

	orderCreate := &models.Orders{
		OrderID:     input.OrderID,
		UserID:      input.UserID,
		TotalAmount: input.TotalAmount,
		Status:      constants.Pending,
	}
	if err := order_repo.CreateOrder(orderCreate); err != nil {
		log.Error().Err(err).Msgf("CreateOrderService-> Error creating order for orderID: %s", input.OrderID)
		return err
	}

	orderItems := make([]*models.OrderItems, 0)
	for _, v := range input.ItemsIDs {
		orderItems = append(orderItems, &models.OrderItems{
			OrderID: orderCreate.OrderID,
			ItemId:  v,
		})
	}

	if err := order_items_repo.CreateOrderItems(orderItems); err != nil {
		log.Error().Err(err).Msgf("CreateOrderService-> Error creating order items for orderID: %s", input.OrderID)
		return err
	}

	orderQueue <- order

	return nil

}

func GetOrderStatus(orderID string) (string, error) {
	orderStatus, err := order_repo.GetOrderStatus(orderID)
	if err != nil {
		log.Error().Err(err).Msgf("GetOrderStatus-> Error getting order status for orderID: %s", orderID)
		return "", err
	}

	return orderStatus, nil

}

func GetOrderMetric() (*response.OrderMetrics, error) {

	totalOrderProcessed, err := order_repo.TotalOrderProcessed()
	if err != nil {
		log.Error().Err(err).Msg("GetOrderMetric-> order_repo.TotalOrderProcessed-> error in getting total order processed")
		return nil, err
	}

	avgOrderProcessingTime, err := order_repo.AvgOrderProcessingTime()
	if err != nil {
		log.Error().Err(err).Msg("GetOrderMetric-> order_repo.AvgOrderProcessingTime-> error in getting avg order processing time")
		return nil, err
	}

	orderStatusByCount, err := order_repo.OrderStatusByCount()
	if err != nil {
		log.Error().Err(err).Msg("GetOrderMetric-> order_repo.OrderStatusByCount-> error in getting  order status by count")
		return nil, err
	}

	response := &response.OrderMetrics{
		TotalOrders:       totalOrderProcessed,
		AvgProcessingTime: avgOrderProcessingTime,
		OrdersByStatus:    orderStatusByCount,
	}

	return response, err
}
