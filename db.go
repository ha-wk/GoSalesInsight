//HERE CONCERNS CAN BE separated in another wrapper db insatance

package main

import (
	"database/sql"
	"encoding/csv"
	"os"
	"strconv"
	//"time"
	"io"
	"log"
)

func InitDB() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=yourpassword dbname=sales_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func logToFile(filename, message string) {
	f, err := os.OpenFile("./logs/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open log file: %v", err)
		return
	}
	defer f.Close()
	if _, err := f.WriteString(message + "\n"); err != nil {    //TODO
		log.Printf("Failed to write to log file: %v", err)
	}
}

func LoadCSVData(db *sql.DB, csvPath string) error {
	// Clear existing data
	_, err := db.Exec(`
		DELETE FROM orders;
		DELETE FROM customers;
		DELETE FROM products;
		DELETE FROM categories;
		DELETE FROM regions;
	`)
	if err != nil {
		return err
	}

	file, err := os.Open(csvPath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip header
	if err != nil {
		return err
	}

	// Maps to avoid duplicate entries
	customers := make(map[string]bool)
	products := make(map[string]bool)
	categories := make(map[string]bool)
	regions := make(map[string]bool)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logToFile("refresh.log", "CSV read error: "+err.Error())
			continue
		}

		// Parse CSV fields
		orderID := record[0]
		productID := record[1]
		customerID := record[2]
		productName := record[3]
		categoryName := record[4]
		regionName := record[5]
		dateOfSale := record[6]
		quantity, _ := strconv.Atoi(record[7])
		unitPrice, _ := strconv.ParseFloat(record[8], 64)
		discount, _ := strconv.ParseFloat(record[9], 64)
		shippingCost, _ := strconv.ParseFloat(record[10], 64)
		paymentMethod := record[11]
		customerName := record[12]
		customerEmail := record[13]         //TODO
		customerAddress := record[14]

		// Insert Category
		if !categories[categoryName] {
			_, err = db.Exec("INSERT INTO categories (category_id, category_name) VALUES ($1, $2)", "CAT"+categoryName[:3], categoryName)
			if err != nil {
				logToFile("refresh.log", "Category insert error: "+err.Error())
				continue
			}
			categories[categoryName] = true
		}

		// Insert Region
		if !regions[regionName] {
			_, err = db.Exec("INSERT INTO regions (region_id, region_name) VALUES ($1, $2)", "REG"+regionName[:3], regionName)
			if err != nil {
				logToFile("refresh.log", "Region insert error: "+err.Error())
				continue
			}
			regions[regionName] = true
		}

		// Insert Customer
		if !customers[customerID] {
			_, err = db.Exec("INSERT INTO customers (customer_id, customer_name, email, address) VALUES ($1, $2, $3, $4)",
				customerID, customerName, customerEmail, customerAddress)
			if err != nil {
				logToFile("refresh.log", "Customer insert error: "+err.Error())
				continue
			}
			customers[customerID] = true
		}

		// Insert Product
		if !products[productID] {
			_, err = db.Exec("INSERT INTO products (product_id, product_name, category_id) VALUES ($1, $2, $3)",
				productID, productName, "CAT"+categoryName[:3])
			if err != nil {
				logToFile("refresh.log", "Product insert error: "+err.Error())
				continue
			}
			products[productID] = true
		}

		// Insert Order
		_, err = db.Exec(`
			INSERT INTO orders (order_id, product_id, customer_id, region_id, date_of_sale, quantity_sold, unit_price, discount, shipping_cost, payment_method)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
			orderID, productID, customerID, "REG"+regionName[:3], dateOfSale, quantity, unitPrice, discount, shippingCost, paymentMethod)
		if err != nil {
			logToFile("refresh.log", "Order insert error: "+err.Error())
			continue
		}
	}
	return nil
}