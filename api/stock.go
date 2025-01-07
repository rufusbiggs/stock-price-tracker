package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const baseURL = "https://www.alphavantage.co/query"

// Fetches the stock data and returns the parsed closing price and timestamp
func FetchStockPrice(symbol string, apiKey string) (string, float64, error) {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"function": "TIME_SERIES_INTRADAY",
			"symbol": symbol,
			"interval": "1min",
			"apikey": apiKey,
		}).
		Get(baseURL)

	if err != nil {
		return "", 0, err
	}

	if resp.IsError() {
		return "", 0, fmt.Errorf("error: %s", resp.Status())
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", 0, err
	}

	// Extract timeseries data
	timeSeries, ok := result["Time Series (1min)"].(map[string]interface{})
	if !ok {
		return "", 0, fmt.Errorf("failed to parse timeseries")
	}

	var latestTimestamp string
	var latestPrice float64
	for timestamp, data := range timeSeries {
		latestTimestamp = timestamp
		priceData := data.(map[string]interface{})
		closePrice := priceData["4. close"].(string)

		// convert to float
		latestPrice = parsedPrice(closePrice)

		break // only need the latest price
	}
	
	return latestTimestamp, latestPrice, nil
}

func parsedPrice(parsedStr string) float64 {
	var price float64
	fmt.Sscanf(parsedStr, "%f", &price)
	return price
}

