package main

import (
	"database/sql"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type RevenueResponse struct {
	Total float64 `json:"total"`       //TOTAL COUNT
}

type RevenueByGroup struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`          //REVENUE FEILDS
	Total float64 `json:"total"`
}

func RefreshData(c *gin.Context, db *sql.DB) {
	err := LoadCSVData(db, "./data/sales_data.csv")
	if err != nil {
		logToFile("refresh.log", "Failed to refresh data: "+err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh data"})    //WE WILL REFRESH DATA VIA CSV
		return
	}
	logToFile("refresh.log", "Data refresh successful at "+time.Now().String())
	c.JSON(http.StatusOK, gin.H{"message": "Data refreshed successfully"})
}

// TODO , MAKE QUERIES IN WRAPPER CLASS
func GetTotalRevenue(c *gin.Context, db *sql.DB) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	query := `
		SELECT COALESCE(SUM(quantity_sold * (unit_price * (1 - discount))), 0) as total
		FROM orders
		WHERE date_of_sale BETWEEN $1 AND $2                                 //DB QUERY,CAN BE PUT INTO CONSATNATS
	`
	var total float64
	err := db.QueryRow(query, startDate, endDate).Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate revenue"})
		return
	}
	c.JSON(http.StatusOK, RevenueResponse{Total: total})
}

//REVENUE BY PRODUCT WE WANT TO CHECK
func GetRevenueByProduct(c *gin.Context, db *sql.DB) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	query := `
		SELECT p.product_id, p.product_name, COALESCE(SUM(o.quantity_sold * (o.unit_price * (1 - o.discount))), 0) as total
		FROM orders o
		JOIN products p ON o.product_id = p.product_id
		WHERE o.date_of_sale BETWEEN $1 AND $2
		GROUP BY p.product_id, p.product_name
	`
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate revenue by product"})
		return
	}
	defer rows.Close()

	var results []RevenueByGroup
	for rows.Next() {
		var r RevenueByGroup
		if err := rows.Scan(&r.ID, &r.Name, &r.Total); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan results"})
			return
		}
		results = append(results, r)
	}
	c.JSON(http.StatusOK, results)
}

//BY CATEGORY IF WE WANT TO CHECK
func GetRevenueByCategory(c *gin.Context, db *sql.DB) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	query := `
		SELECT c.category_id, c.category_name, COALESCE(SUM(o.quantity_sold * (o.unit_price * (1 - o.discount))), 0) as total
		FROM orders o
		JOIN products p ON o.product_id = p.product_id
		JOIN categories c ON p.category_id = c.category_id
		WHERE o.date_of_sale BETWEEN $1 AND $2
		GROUP BY c.category_id, c.category_name
	`
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate revenue by category"})
		return
	}
	defer rows.Close()

	var results []RevenueByGroup
	for rows.Next() {
		var r RevenueByGroup
		if err := rows.Scan(&r.ID, &r.Name, &r.Total); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan results"})
			return
		}
		results = append(results, r)
	}
	c.JSON(http.StatusOK, results)
}

func GetRevenueByRegion(c *gin.Context, db *sql.DB) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	query := `
		SELECT r.region_id, r.region_name, COALESCE(SUM(o.quantity_sold * (o.unit_price * (1 - o.discount))), 0) as total
		FROM orders o
		JOIN regions r ON o.region_id = r.region_id
		WHERE o.date_of_sale BETWEEN $1 AND $2
		GROUP BY r.region_id, r.region_name
	`
	rows, err := db.Query(query, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate revenue by region"})
		return
	}
	defer rows.Close()

	var results []RevenueByGroup
	for rows.Next() {
		var r RevenueByGroup
		if err := rows.Scan(&r.ID, &r.Name, &r.Total); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan results"})
			return
		}
		results = append(results, r)
	}
	c.JSON(http.StatusOK, results)
}