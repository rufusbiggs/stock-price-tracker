package main

import (
	"fmt"
	"log"
	"stock-price-tracker/api"
	"stock-price-tracker/db"
)

func main() {
	connStr := "postgres://rufusbiggs:Curry123!@localhost/stock_tracker?sslmode=disable"
	db.InitDB(connStr)

	apiKey := "WO363FDOPGSZ33EN"
	symbol := "AAPL" // for testing fetch Apple stock prices

	timestamp, price, err := api.FetchStockPrice(symbol, apiKey)
	if err != nil {
		log.Fatalf("Error fetching stock data: %v", err)
	}

	err = db.SaveStockData(symbol, price, timestamp)
	if err != nil {
		log.Fatalf("Error saving stock data: %v", err)
	}

	fmt.Println("Stock Data saved successfully!")
}