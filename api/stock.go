package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

const baseURL = "https://www.alphavantage.co/query"

func FetchStockPrice(symbol string, apiKey string) (map[string]interface{}, error) {
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
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("error: %s", resp.Status())
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, err
	}
	
	return result, nil
}

