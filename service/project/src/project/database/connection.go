package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //for psql
	"project/config"
)

// NewConnection create new database connection
func NewConnection(cfg config.AppConfig) *sql.DB {
	pattern := "user=%s password=%s dbname=%s host=%s sslmode=disable"

	dbInfo := fmt.Sprintf(pattern, cfg.GetDatabaseUser(), cfg.GetDatabasePassword(), cfg.GetDatabaseName(), cfg.GetDatabaseHost())
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	return db
}
