package main

import (
	"fmt"
	"log"
	"stock-price-tracker/api"
)

func main() {
	apiKey := "WO363FDOPGSZ33EN"
	symbol := "AAPL" // for testing fetch Apple stock prices

	stockData, err := api.FetchStockPrice(symbol, apiKey)
	if err != nil {
		log.Fatalf("Error fetching stock data: %v", err)
	}

	fmt.Printf("Stock Data: %v\n", stockData)
}