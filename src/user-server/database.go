package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func (config Config) CreateConnection() *sql.DB {
	DB, err := sql.Open(config.DBType, config.DBHost)
	if err != nil {
		log.Fatal(err)
	}
	//don't forget to close the db
	// defer DB.Close()
	return DB
}
