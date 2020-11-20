package helpers

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (db *sql.DB) {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	return db

}
