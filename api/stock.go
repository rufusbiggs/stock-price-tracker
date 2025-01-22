package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
	"net/http"
	"os"
)

type TimestampValue struct {
	Timestamp string
	Price float64
}

const baseURL = "https://www.alphavantage.co/query"

// Fetches the stock data and returns the parsed closing price and timestamp
func FetchStockPrice(symbol string, apiKey string, client *http.Client) (string, float64, []TimestampValue, error) {
	// Initialize the resty client with the provided http.Client
	restyClient := resty.New().
		SetTimeout(10 * time.Second).
		SetRetryCount(3).
    	SetRetryWaitTime(2 * time.Second).
    	SetRetryMaxWaitTime(10 * time.Second)

	// Use the provided client for the request
	resp, err := restyClient.R().
		SetQueryParams(map[string]string{
			"function": "TIME_SERIES_INTRADAY",
			"symbol":   symbol,
			"interval": "60min",
			"apikey":   apiKey,
		}).
		Get(baseURL)

	if err != nil {
		if os.IsTimeout(err) {
			log.Printf("API request timed out: %v", err)
			return "", 0, nil, fmt.Errorf("API timeout: %v", err)
		}
		return "", 0, nil, fmt.Errorf("error making API request: %v", err)
	}

	if resp.IsError() {
		// Log the error response for debugging purposes
		log.Printf("API Request failed with status: %s and response: %s", resp.Status(), resp.String())
		return "", 0, nil, fmt.Errorf("API error: %s", resp.Status())
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", 0, nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	// Extract timeseries data
	timeSeries, ok := result["Time Series (60min)"].(map[string]map[string]string)
	if !ok {
		return "", 0, nil, fmt.Errorf("failed to parse timeseries data from response")
	}

	// Get the latest data
	var latestTimestamp string
	var latestPrice float64

	// iterate through keys to find latest timestamp
	for timestamp := range timeSeries {
		if latestTimestamp == "" || timestamp > latestTimestamp {
			latestTimestamp = timestamp
		}
	}

	if latestTimestamp == "" {
		return "", 0, nil, fmt.Errorf("failed to extract latest timestamp")
	}

	// got the latest date, now going to filter only datapoints on this day
	var daysPrices []TimestampValue
	latestDayStr := latestTimestamp[:10]

	for timestamp, data := range timeSeries {
		if timestamp[:10] == latestDayStr {
			closePriceStr, ok := data["4. close"]
			if !ok {
				fmt.Printf("failed to parse close price data")
				continue
			}
			price := parsedPrice(closePriceStr)
			daysPrices = append(daysPrices, TimestampValue{
				Timestamp: timestamp,
				Price:     price,
			})
		}
	}

	latestPrice = daysPrices[0].Price

	return latestTimestamp, latestPrice, daysPrices, nil
}

// Helper function to parse the price string into a float
func parsedPrice(parsedStr string) float64 {
	var price float64
	fmt.Sscanf(parsedStr, "%f", &price)
	return price
}
