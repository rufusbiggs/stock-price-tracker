package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"stock-price-tracker/db"
)

func StartServer() {
	router := gin.Default()

	router.GET("/stocks/:symbol/latest", getLatestStockPrice)
	router.GET("/stocks/:symbol/history", getHistoricalStockPrices)

	router.Run(":8080")
}

func getLatestStockPrice(c *gin.Context) {
	symbol := c.Param("symbol")
	latest, err := db.GetLatestStockPrice(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, latest)
}

func getHistoricalStockPrices(c *gin.Context) {
	symbol := c.Param("symbol")
	historicalData, err := db.GetHistoricalStockPrices(symbol)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, historicalData)
}


// "Meta Data": {
//         "1. Information": "Intraday (1min) open, high, low, close prices and volume",
//         "2. Symbol": "AAPL",
//         "3. Last Refreshed": "2025-01-16 19:59:00",
//         "4. Interval": "1min",
//         "5. Output Size": "Compact",
//         "6. Time Zone": "US/Eastern"
//     },
//     "Time Series (1min)": {
//         "2025-01-16 19:59:00": {
//             "1. open": "228.4500",
//             "2. high": "228.4500",
//             "3. low": "228.3500",
//             "4. close": "228.3501",
//             "5. volume": "6247"