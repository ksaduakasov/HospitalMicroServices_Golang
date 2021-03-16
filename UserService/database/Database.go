package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func OpenDB(dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Println("Connection for database couldn't be established")
		return nil, err
	}
	return pool, nil
}
