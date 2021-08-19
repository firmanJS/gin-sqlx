package database

import (
	"fmt"
	"log"

	"github.com/firmanJS/gin-sqlx/src/config"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 100
	maxIdleConns    = 30
	Version         = "1.0.0"
)

func NewConnection() *sqlx.DB {
	var err error

	config, errs := config.AppConfig()

	if errs != nil {
		log.Fatal(err)
	}

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DB.Host, config.DB.Port, config.DB.Username, config.DB.Password, config.DB.Name, "disable")

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		log.Panic("[CONFIG DB] error connect :", err)
	}

	if err = db.Ping(); err != nil {
		log.Panic("[CONFIG DB] error ping :", err)
	}

	db.SetMaxOpenConns(maxOpenConns)       // The default is 0 (unlimited)
	db.SetMaxIdleConns(maxIdleConns)       // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(connMaxLifetime) // 0, connections are reused forever.
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	return db
}
