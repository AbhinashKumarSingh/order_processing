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