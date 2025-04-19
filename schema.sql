CREATE DATABASE sales_db;

\c sales_db;

CREATE TABLE customers (
    customer_id VARCHAR(50) PRIMARY KEY,
    customer_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    address TEXT NOT NULL
);

CREATE TABLE categories (
    category_id VARCHAR(50) PRIMARY KEY,
    category_name VARCHAR(100) NOT NULL
);

CREATE TABLE regions (
    region_id VARCHAR(50) PRIMARY KEY,
    region_name VARCHAR(100) NOT NULL
);

CREATE TABLE products (
    product_id VARCHAR(50) PRIMARY KEY,
    product_name VARCHAR(100) NOT NULL,
    category_id VARCHAR(50) REFERENCES categories(category_id)
);

CREATE TABLE orders (
    order_id VARCHAR(50) PRIMARY KEY,
    product_id VARCHAR(50) REFERENCES products(product_id),
    customer_id VARCHAR(50) REFERENCES customers(customer_id),
    region_id VARCHAR(50) REFERENCES regions(region_id),
    date_of_sale DATE NOT NULL,
    quantity_sold INTEGER NOT NULL,
    unit_price FLOAT NOT NULL,
    discount FLOAT NOT NULL,
    shipping_cost FLOAT NOT NULL,
    payment_method VARCHAR(50) NOT NULL
);