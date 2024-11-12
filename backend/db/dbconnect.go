package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func ConnectDB() (*pgxpool.Pool, error) {

	if pool != nil {
		return pool, nil
	}

	dsn := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Printf("Unable to parse database config: %v\n", err)
		return nil, err
	}

	config.MaxConns = 10
	config.MinConns = 2

	// Create connection pool
	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		log.Printf("Unable to ping database: %v\n", err)
		return nil, err
	}

	log.Println("Successfully connected to database!")
	return pool, nil
}
