package db

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/Babushkin05/simple-marketplace/goods-service/internal/config"
)

//go:embed schema.sql
var schema string

func MustInitPostgres(cfg config.Config) *sqlx.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}
	log.Println("Connected to PostgreSQL")

	if _, err := db.Exec(schema); err != nil {
		log.Fatalf("failed to apply schema: %v", err)
	}
	log.Println("DB schema applied")

	return db
}
