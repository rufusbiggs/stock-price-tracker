package db

import (
	"database/sql"
	_"github.com/lib/pq"
	"log"
)

const createStockTable = `
	CREATE TABLE IF NOT EXISTS stocks (
		id SERIAL PRIMARY KEY,
		symbol VARCHAR(10) NOT NULL,
		price NUMERIC(10, 2) NOT NULL,
		timestamp TIMESTAMP NOT NULL
	);
`

var DB *sql.DB

func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	_, err = DB.Exec(createStockTable)
	if err != nil {
		log.Fatalf("Error creating stock table: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Database connected!")
}

func SaveStockData(symbol string, price float64, timestamp string) error {
	query := "INSERT INTO stocks (symbol, price, timestamp) VALUES ($1, $2, $3)"
	_, err := DB.Exec(query, symbol, price, timestamp)
	return err
}

func GetLatestStockPrice(symbol string) (map[string]interface{}, error) {
	query := "SELECT price, timestamp FROM stocks WHERE symbol = $1 ORDER BY timestamp DESC LIMIT 1"
	row := DB.QueryRow(query, symbol)

	var price float64
	var timestamp string
	err := row.Scan(&price, &timestamp); 
	if err != nil {
		log.Println("Error fetching stock data")
		return nil, err
	}

	return map[string]interface{} {
		"symbol": symbol,
		"price": price,
		"timestamp": timestamp,
	}, nil
}

func GetHistoricalStockPrices(symbol string) ([]map[string]interface{}, error) {
	query := "SELECT price, timestamp FROM stocks WHERE symbol = $1 ORDER BY timestamp DESC LIMIT 100"
	rows, err := DB.Query(query, symbol)
	if err != nil {
		log.Println("Error fetching historical stock data")
		return nil, err
	}
	defer rows.Close()

	var history []map[string]interface{}
	for rows.Next() {
		var price float64
		var timestamp string
		err := rows.Scan(&price, &timestamp)
		if err != nil {
			return nil, err
		}
		history = append(history, map[string]interface{}{
			"price": price,
			"timestamp": timestamp,
		})
	}
	
	return history, nil
} 