package db

import (
	"database/sql"
	_"github.com/lib/pq"
	"log"
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
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