package main

import (
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Initializing database
	db, err := InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	//setting gin for sending requests
	r := gin.Default()

	r.POST("/refresh", func(c *gin.Context) { RefreshData(c, db) })
	r.GET("/revenue/total", func(c *gin.Context) { GetTotalRevenue(c, db) })
	r.GET("/revenue/product", func(c *gin.Context) { GetRevenueByProduct(c, db) })
	r.GET("/revenue/category", func(c *gin.Context) { GetRevenueByCategory(c, db) })
	r.GET("/revenue/region", func(c *gin.Context) { GetRevenueByRegion(c, db) })

	// hitting server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("SerVer fAILED: %v", err)
	}
}