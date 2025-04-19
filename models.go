package main

type Customer struct {
	CustomerID   string `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
}

type Product struct {
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	CategoryID  string `json:"category_id"`
}

type Category struct {
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type noveltyRegion struct {
	RegionID   string `json:"region_id"`
	RegionName string `json:"region_name"`
}

type Order struct {
	OrderID       string  `json:"order_id"`
	ProductID     string  `json:"product_id"`
	CustomerID    string  `json:"customer_id"`
	RegionID      string  `json:"region_id"`
	DateOfSale    string  `json:"date_of_sale"`
	QuantitySold  int     `json:"quantity_sold"`
	UnitPrice     float64 `json:"unit_price"`
	Discount      float64 `json:"discount"`
	ShippingCost  float64 `json:"shipping_cost"`
	PaymentMethod string  `json:"payment_method"`
}