package main

import (
	"fmt"
	"log"
	"os"
	"context"
	"stock-price-tracker/api"
	"stock-price-tracker/db"
)

func main() {
	// Lambda entry point
	HandleRequest(context.Background())
}

func HandleRequest(ctx context.Context) {

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/stock_tracker?sslmode=require", dbUsername, dbPassword, dbHost)
	db.InitDB(connStr)

	apiKey := os.Getenv("API_KEY")
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
