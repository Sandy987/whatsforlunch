package main

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Used for postgres driver support
)

var DB *sqlx.DB

// InitDb fetches the singleton instance of the domain
func InitDb() *sqlx.DB {
	if DB == nil {
		conString := os.Getenv("POSTGRES_URL")

		if conString == "" {
			log.Fatal("POSTGRES_URL environment variable not specified, could not open DB connection")
		}

		newdb, err := sqlx.Connect("postgres", conString)
		if err != nil {
			log.Fatal(err)
		}
		DB = newdb
	}
	return DB
}
