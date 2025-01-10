package main

import (
	"fmt"
	"log"
	"os"
	"context"
	"time"
	"stock-price-tracker/api"
	"stock-price-tracker/db"
	"net/http"
)

func main() {
	// Lambda entry point
	HandleRequest(context.Background())
}

func HandleRequest(ctx context.Context) {

	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/stocks?sslmode=require", dbUsername, dbPassword, dbHost)
	db.InitDB(connStr)

	apiKey := os.Getenv("API_KEY")
	symbol := "AAPL" // for testing fetch Apple stock prices

	// Add timeout for the API request
	client := &http.Client{
		Timeout: 5 * time.Second, // Set timeout to 5 seconds
	}

	timestamp, price, err := api.FetchStockPrice(symbol, apiKey, client)
	if err != nil {
		log.Fatalf("Error fetching stock data: %v", err)
		return
	}

	log.Printf("Fetched stock data: Symbol: %s, Price: %f, Timestamp: %s", symbol, price, timestamp)

	err = db.SaveStockData(symbol, price, timestamp)
	if err != nil {
		log.Fatalf("Error saving stock data: %v", err)
		return
	}

	log.Println("Stock Data saved successfully!")
}
