package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq" // postgres driver
	"github.com/mariosker/taskfrenzy/config"
)

func main() {
	cfg, err := pgx.ParseConfig(config.Envs.DBConnString)
	if err != nil {
		log.Fatal(err)
	}

	// Use pgx stdlib to open a database connection compatible with *sql.DB
	db := stdlib.OpenDB(*cfg)
	defer db.Close()

	// Initialize the storage by pinging the database
	initStorage(db)

	// Set up the PostgreSQL migration driver
	driver, err := postgresMigrate.WithInstance(db, &postgresMigrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new migration instance with PostgreSQL
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", // Your migrations folder
		cfg.Database,                    // Name of the database system
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Log the current migration version
	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	// Command to migrate up or down
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}

func initStorage(db *sql.DB) {
	ctx := context.Background()

	// Ping the database to ensure connection is valid
	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("DB: Failed to connect:", err)
	}

	log.Println("DB: Successfully connected!")
}
