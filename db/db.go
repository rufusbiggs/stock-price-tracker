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