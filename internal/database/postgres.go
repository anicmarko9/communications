package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"communications/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.Config) *pgxpool.Pool {

	connectionString := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
		cfg.DatabaseSSL,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connectionString)

	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping the database: %v", err)
	}

	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DatabaseUser,
		cfg.DatabasePassword,
		cfg.DatabaseHost,
		cfg.DatabasePort,
		cfg.DatabaseName,
		cfg.DatabaseSSL,
	)

	migration, err := migrate.New("file://migrations", databaseURL)

	if err != nil {
		log.Fatalf("Unable to initiate the SQL migrations: %v", err)
	}
	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Unable to run the SQL migrations: %v", err)
	}

	fmt.Println("SQL Migrations applied successfully.")

	return pool
}
