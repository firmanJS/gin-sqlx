package database

import (
	"fmt"
	"log"

	"github.com/firmanJS/gin-sqlx/src/config"
	"github.com/jmoiron/sqlx"
)

func NewConnection(config *config.DB) *sqlx.DB {
	var err error

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.dbHost, config.dbPort, config.dbUsername, config.dbPassword, config.dbName)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Panic("[CONFIG DB] error connect :", err)
	}

	if err = db.Ping(); err != nil {
		log.Panic("[CONFIG DB] error ping :", err)
	}

	return db
}
