package order_service

import (
	"sync"
	"time"

	"example.com/m/v2/constants"
	"example.com/m/v2/db/order_repo"
	"github.com/phuslu/log"
)

type Order struct {
	ID     string `db:"id"`
	Status string `db:"status"`
}

var orderQueue = make(chan *Order, 1000) // A channel acting as the queue
var orders = make(map[string]*Order)     // Map to track orders by ID
var mu sync.Mutex

func Worker() {
	for order := range orderQueue {
		processOrder(order)
	}
}

func processOrder(order *Order) {
	// Step 1: Update the order status to Processing in the database
	order.Status = constants.Processing
	mu.Lock()
	orders[order.ID] = order
	mu.Unlock()

	// Update order status in the database
	updates := map[string]interface{}{
		"status": order.Status,
	}

	if err := order_repo.UpdateStatus(order.ID, updates); err != nil {
		log.Error().Err(err).Msgf("Error updating status in DB for orderID:%s to processing: ", order.ID)
	}

	// Step 3: Simulate time delay for processing (e.g., payment processing, inventory check)
	time.Sleep(5 * time.Second) // Simulate some processing time

	// Step 4: Once processing is completed, update the order status to Completed
	order.Status = constants.Completed
	mu.Lock()
	orders[order.ID] = order
	mu.Unlock()
	updates["status"] = constants.Completed
	updates["completed_at"] = time.Now()

	if err := order_repo.UpdateStatus(order.ID, updates); err != nil {
		log.Error().Err(err).Msgf("Error updating status in DB for orderID:%s to completed: ", order.ID)
	}
}
