package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgreSQLStorage(cfg pgx.ConnConfig) (*pgxpool.Pool, error) {
	// Create the connection string from the provided config
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	// Create a context (optional, can be used for connection timeout)
	ctx := context.Background()

	// Initialize the connection pool
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	return pool, nil
}
