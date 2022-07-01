CREATE orders_by;

CREATE table items (
    item_id SERIAL PRIMARY KEY,
    item_code VARCHAR(50) NOT NULL,
    description VARCHAR(50) NOT NULL,
    quantity INT,
    order_id INT
);

CREATE table orders (
    order_id SERIAL PRIMARY KEY,
    customer_name VARCHAR(30) NOT NULL,
    ordered_at VARCHAR(20) NOT NULL
);