package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //for psql
	"gopkg.in/gorp.v2"
	"log"
	"project/config"
	"project/model"
)

func mapEntities(dbMap *gorp.DbMap) {
	//TODO: implement DDL "id UUID PRIMARY KEY DEFAULT gen_random_uuid(),"
	dbMap.AddTableWithName(model.Project{}, "project").SetKeys(false, "UUID")
}

// NewConnection create new database connection
func NewConnection(cfg config.AppConfig) *gorp.DbMap {
	pattern := "user=%s password=%s dbname=%s host=%s sslmode=disable"

	dbInfo := fmt.Sprintf(pattern, cfg.GetDatabaseUser(), cfg.GetDatabasePassword(), cfg.GetDatabaseName(), cfg.GetDatabaseHost())
	db, err := sql.Open("postgres", dbInfo)
	checkErr(err, "sql.Open failed")

	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	mapEntities(dbmap)

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
