package psql

import (
	"database/sql"

	"Users/config"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnectionStrings.ServiceDb)

	if err != nil {
		panic(err)
	}

	return db, nil
}
