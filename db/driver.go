package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func init() {
	dsn := "postgres://postgres:postgres@localhost:5432/signoz?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}
}

func Ping() error {
	return db.Ping()
}

func GetDBInstance() *sql.DB {
	return db
}
