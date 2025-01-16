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