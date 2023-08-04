package database

import (
	"SecondAssignment/config"
	"database/sql"
	"log"
)

func Open(cfg config.Schema) (*sql.DB, error) {
	db, err := sql.Open(cfg.SQL.Driver, cfg.SQL.DataSourceName)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	return db, err
}
