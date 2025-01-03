package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"sync"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDbConnection() *sql.DB {
	once.Do(
		func() {
			connStr := "user=postgres dbname=postgres sslmode=disable"
			var err error
			db, err = sql.Open("postgres", connStr)
			if err != nil {
				panic(fmt.Sprintf("Failed to open database: %v", err))
			}

			err = db.Ping()
			if err != nil {
				panic(fmt.Sprintf("Failed to ping database: %v", err))
			}

			fmt.Println("Successfully connected!")
		},
	)

	return db
}
