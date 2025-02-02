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
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

func main() {
	// Check run mode to run locally or on Lambda
	runMode := os.Getenv("RUN_MODE")
	if runMode == "local" {
		log.Printf("Running locally. Starting server...")
		// Load environment variables for security
		err := godotenv.Load()
        if err != nil {
            log.Fatalf("Error loading .env file: %v", err)
        }
        log.Println("Loaded environment variables from .env")
		initDatabase()
		// Set default values for local testing
		os.Setenv("_LAMBDA_SERVER_PORT", "8080")
		os.Setenv("AWS_LAMBDA_RUNTIME_API", "localhost:8081")
	} else {
		// Lambda entry point
		initDatabase()
		log.Printf("Running in AWS Lambda. Starting Lambda handler...")
		lambda.Start(HandleRequest)
	}
	api.StartServer()
}

func initDatabase() {
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	connStr := fmt.Sprintf("postgres://%s:%s@%s/stocks?sslmode=require", dbUsername, dbPassword, dbHost)
	db.InitDB(connStr)
}

func HandleRequest(ctx context.Context) error {

	apiKey := os.Getenv("API_KEY")
	symbol := "AAPL" // for testing fetch Apple stock prices

	// Add timeout for the API request
	client := &http.Client{
		Timeout: 15 * time.Second, // Set timeout to 15 seconds
	}

	log.Println("Making API request...")
	timestamp, price, daysPrices, err := api.FetchStockPrice(symbol, apiKey, client)
	if err != nil {
		log.Fatalf("Error fetching stock data: %v", err)
		return err
	}

	log.Printf("Latest stock data: Symbol: %s, Price: %f, Timestamp: %s", symbol, price, timestamp)

	for _, data := range daysPrices {
		daysPrice := data.Price
		daysTimestamp := data.Timestamp

		err = db.SaveStockData(symbol, daysPrice, daysTimestamp)
		if err != nil {
			log.Fatalf("Error saving stock data: %v", err)
			return err
		}
	}

	log.Println("Stock Data saved successfully!")
	return nil
}
