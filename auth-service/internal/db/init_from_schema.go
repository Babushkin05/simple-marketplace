package db

import (
	"database/sql"
	_ "embed"
	"fmt"
)

//go:embed schema.sql
var schema string

func InitSchemaFromFile(db *sql.DB) error {
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to execute schema: %w", err)
	}
	return nil
}
