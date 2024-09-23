package psql

import (
	"fmt"

	"Users/config"

	"database/sql"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnectionStrings.ServiceDb)

	if err != nil {
		return nil, fmt.Errorf("database connecting execution error: %v", err)
	}

	return db, nil
}
