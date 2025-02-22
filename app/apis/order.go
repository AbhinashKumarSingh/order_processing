package apis

import (
	"net/http"

	"example.com/m/v2/requests"
	"example.com/m/v2/service/order_service"
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
)

func OrderCreate(c echo.Context) error {

	var orderInput requests.OrderReq
	if err := c.Bind(&orderInput); err != nil {
		log.Error().Err(err).Msg("OrderCreate: Failed to bind request")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := order_service.CreateOrderService(orderInput); err != nil {
		log.Error().Err(err).Msg("OrderCreate: Failed to create order")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not create order"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Order created successfully"})
}

func GetOrderStatus(c echo.Context) error {
	orderID := c.QueryParam("order_id")

	if orderID == "" {
		log.Error().Msg("GetOrderStatus: orderID is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	orderStatus, err := order_service.GetOrderStatus(orderID)
	if err != nil {
		log.Error().Err(err).Msg("GetOrderStatus: Failed to get order")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get order status"})
	}
	return c.JSON(http.StatusOK, map[string]string{"order_status": orderStatus})

}

func GetOrderMetrics(c echo.Context) error {

	response, err := order_service.GetOrderMetric()
	if err != nil {
		log.Error().Err(err).Msg("GetOrderStatus: Failed to get order")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get order status"})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"order_metrics": response})

}
