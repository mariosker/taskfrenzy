package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mariosker/taskfrenzy/cmd/api"
	"github.com/mariosker/taskfrenzy/config"
	"github.com/mariosker/taskfrenzy/db"
)

func main() {
	// Create the connection configuration for PostgreSQL
	cfg, err := pgx.ParseConfig(config.Envs.DBConnString)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the PostgreSQL connection pool
	dbPool, err := db.NewPostgreSQLStorage(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	// Initialize the storage by pinging the database
	initStorage(dbPool)

	// Start the API server
	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), dbPool)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

// initStorage verifies the connection to the database
func initStorage(dbPool *pgxpool.Pool) {
	ctx := context.Background()

	// Ping the database to ensure connection is valid
	err := dbPool.Ping(ctx)
	if err != nil {
		log.Fatal("DB: Failed to connect:", err)
	}

	log.Println("DB: Successfully connected!")
}
