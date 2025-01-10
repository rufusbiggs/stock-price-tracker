package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
	"net/http"
)

const baseURL = "https://www.alphavantage.co/query"

// Fetches the stock data and returns the parsed closing price and timestamp
func FetchStockPrice(symbol string, apiKey string, client *http.Client) (string, float64, error) {
	// Initialize the resty client with the provided http.Client
	restyClient := resty.New()
	restyClient.SetTimeout(10 * time.Second) // Set a timeout for the request (if desired)

	// Use the provided client for the request
	resp, err := restyClient.R().
		SetQueryParams(map[string]string{
			"function": "TIME_SERIES_INTRADAY",
			"symbol":   symbol,
			"interval": "1min",
			"apikey":   apiKey,
		}).
		Get(baseURL)

	if err != nil {
		return "", 0, fmt.Errorf("error making API request: %v", err)
	}

	if resp.IsError() {
		// Log the error response for debugging purposes
		log.Printf("API Request failed with status: %s and response: %s", resp.Status(), resp.String())
		return "", 0, fmt.Errorf("API error: %s", resp.Status())
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", 0, fmt.Errorf("Error unmarshaling response: %v", err)
	}

	// Extract timeseries data
	timeSeries, ok := result["Time Series (1min)"].(map[string]interface{})
	if !ok {
		return "", 0, fmt.Errorf("Failed to parse timeseries data from response")
	}

	// Get the latest data
	var latestTimestamp string
	var latestPrice float64
	for timestamp, data := range timeSeries {
		latestTimestamp = timestamp
		priceData := data.(map[string]interface{})
		closePrice := priceData["4. close"].(string)

		// Convert to float
		latestPrice = parsedPrice(closePrice)

		break // Only need the latest price
	}

	return latestTimestamp, latestPrice, nil
}

// Helper function to parse the price string into a float
func parsedPrice(parsedStr string) float64 {
	var price float64
	fmt.Sscanf(parsedStr, "%f", &price)
	return price
}
