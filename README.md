# order_processing


# Order Processing Backend System

This repository contains a backend system for order processing, designed to manage and process orders in an e-commerce platform. It demonstrates proficiency in building modular, maintainable, and scalable systems, covering database design, queuing, distributed system fundamentals, and metrics reporting.

## Table of Contents

-   [Objective](#objective)
-   [Problem Statement](#problem-statement)
-   [Technologies Used](#technologies-used)
-   [Setup and Installation](#setup-and-installation)
-   [Running the Application](#running-the-application)
-   [API Endpoints](#api-endpoints)
-   [Database Schema](#database-schema)
-   [Queue Processing](#queue-processing)
-   [Metrics Reporting](#metrics-reporting)
-   [Design Decisions and Trade-offs](#design-decisions-and-trade-offs)
-   [Assumptions](#assumptions)
-   [Future Improvements](#future-improvements)
-   [Contributing](#contributing)
-   [License](#license)

## Objective

The objective of this assignment is to evaluate the candidate’s ability to design and implement a backend system that demonstrates proficiency in building modular, maintainable, and scalable systems while covering database design, queuing, distributed system fundamentals, and metrics reporting.

## Problem Statement

Build a backend system to manage and process orders in an e-commerce platform. The system should:

-   Provide a RESTful API to accept orders.
-   Simulate asynchronous order processing using an in-memory queue.
-   Provide an API to check the status of orders.
-   Implement an API to fetch key metrics.

## Technologies Used

-   **Go (Golang):** For backend development.
-   **MySQL:** For database storage.
-   **GORM:** For ORM (Object-Relational Mapping).
-   **Echo:** For RESTful API framework.
-   **In-memory queue (channels):** For asynchronous processing.
-   **Logrus:** For logging.

## Setup and Installation

1.  **Prerequisites:**
    -   Go 1.16 or later installed.
    -   MySQL installed and running.
    -   Git installed.

2.  **Clone the repository:**

    ```bash
    git clone [https://github.com/AbhinashKumarSingh/order_processing.git]
    cd order-processing
    ```

3.  **Create the MySQL database and tables:**

    -   Create a database named `orders_db`.
    -   Run the provided SQL script (`schema.sql`) to create the `orders` and `order_items` tables.

    ```sql
    CREATE DATABASE orders_db;
    USE orders_db;

    CREATE TABLE orders (
        order_id    VARCHAR(36) PRIMARY KEY,
        user_id     VARCHAR(36) NOT NULL,
        total_amount DECIMAL(10,2) NOT NULL,
        status      ENUM('Pending', 'Processing', 'Completed') NOT NULL DEFAULT 'Pending',
        created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        completed_at TIMESTAMP DEFAULT NULL
    );

    CREATE TABLE order_items (
        order_id VARCHAR(36),
        item_id  VARCHAR(36),
        PRIMARY KEY (order_id, item_id),
        FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE
    );
    ```

4.  **Configure database connection:**
    -   Update the database connection details in `config/config.go`.

    ```go
    dsn := "root:{{password}}@unix(/opt/homebrew/var/mysql/mysql.sock)/orders_db?parseTime=true&tls=false"
    ```
    - Replace the username, password and socket path with your local mysql setup.

5.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

## Running the Application

1.  **Start the application:**

    ```bash
    go run main.go
    ```

    The application will start on port 8000.

## API Endpoints

-   **Create Order (POST /order-service/order/create):**

    ```json
    {
        "order_id": "123e4567-e89b-12d3-a456-426614174000",
        "user_id": "user123",
        "items_ids": ["item1", "item2"],
        "total_amount": 100.00
    }
    ```

    Response:

    ```json
    {
        "message": "Order created successfully"
    }
    ```

-   **Get Order Status (GET /order-service/order/order-status?order_id={order_id}):**

    ```bash
    curl "http://localhost:8000/order-service/order/order-status?order_id=123e4567-e89b-12d3-a456-426614174000"
    ```

    Response:

    ```json
    {
        "order_status": "Completed"
    }
    ```

-   **Get Order Metrics (GET /order-service/order/order-metric):**

    ```bash
    curl "http://localhost:8000/order-service/order/order-metric"
    ```

    Response:

    ```json
    {
        "order_metrics": {
            "total_orders": 10,
            "avg_processing_time": 5,
            "orders_by_status": [
                {
                    "status": "Completed",
                    "count": 10
                }
            ]
        }
    }
    ```

## Database Schema

-   **orders:**
    -   `order_id` (VARCHAR(36), PRIMARY KEY)
    -   `user_id` (VARCHAR(36), NOT NULL)
    -   `total_amount` (DECIMAL(10,2), NOT NULL)
    -   `status` (ENUM('Pending', 'Processing', 'Completed'), NOT NULL, DEFAULT 'Pending')
    -   `created_at` (TIMESTAMP, DEFAULT CURRENT_TIMESTAMP)
    -   `completed_at` (TIMESTAMP, DEFAULT NULL)
-   **order_items:**
    -   `order_id` (VARCHAR(36), FOREIGN KEY)
    -   `item_id` (VARCHAR(36), PRIMARY KEY)

## Queue Processing

-   An in-memory channel (`orderQueue`) is used as a queue.
-   A worker goroutine processes orders from the queue asynchronously.
-   The worker updates the order status in the database to `Processing` and then `Completed` after a simulated processing time.

## Metrics Reporting

-   The `/order-metric` API endpoint fetches and returns the following metrics:
    -   Total number of orders processed.
    -   Average processing time for orders.
    -   Count of orders in each status.

## Design Decisions and Trade-offs

-   **In-memory queue:** Used for simplicity and to meet the assignment requirements. In a production environment, a more robust message queue like RabbitMQ or Kafka would be preferred.
-   **MySQL:** Chosen for its widespread use and ease of setup. Other databases like PostgreSQL or SQLite could also have been used.
-   **GORM:** Used for database interaction to simplify database operations and improve maintainability.
-   **Echo:** Used for its simplicity and performance in handling RESTful APIs.
-   **Simulated processing time:** A `time.Sleep` is used to simulate order processing. In a real-world scenario, this would be replaced with actual processing logic.
-   **Concurrency:** Go's goroutines and channels are used to handle concurrent order processing.

## Assumptions

-   The application assumes that the database connection details are correctly configured.
-   The application assumes that the `order_id` is unique.
-   The application assumes that the items provided in the `items_ids` array are valid.
-   The in-memory queue is sufficient for the simulated load of 1,000 concurrent orders.
-   The processing time simulation of 5 seconds is acceptable.

## Future Improvements

-   Implement a more robust message queue (e.g., RabbitMQ, Kafka,Nats).
-   Add error handling and validation for API requests.
-   Implement proper logging and monitoring.
-   Add unit and integration tests.
-   Implement a load balancer for better scalability.
-   Add authentication and authorization for API endpoints.
-   Implement database connection pooling.
-   Improve the processing time simulation to more closely resemble real-