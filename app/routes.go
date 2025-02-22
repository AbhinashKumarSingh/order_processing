package main

import (
	"example.com/m/v2/apis"
	"example.com/m/v2/service/order_service"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	go order_service.Worker()
	v2 := e.Group("/order-service")
	orderRoutes(v2)
}

func orderRoutes(router *echo.Group) {
	routes := router.Group("/order")
	routes.POST("/create", apis.OrderCreate)
	routes.GET("/order-status", apis.GetOrderStatus)
	routes.GET("/order-metric", apis.GetOrderMetrics)
}
