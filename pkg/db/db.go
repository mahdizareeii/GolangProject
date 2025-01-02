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

const (
	host     = "localhost"
	port     = 5432
	user     = "test_user"
	password = "1234"
	dbname   = "postgres"
)

func GetDbConnection() *sql.DB {
	once.Do(
		func() {
			//var err error
			//db, err = sql.Open("postgres", "jdbc:postgresql://localhost:5432/postgres")
			//if err != nil {
			//	panic(err)
			//}

			psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
				"password=%s dbname=%s sslmode=disable",
				host, port, user, password, dbname)
			db, err := sql.Open("postgres", psqlInfo)
			if err != nil {
				panic(err)
			}
			defer db.Close()

			err = db.Ping()
			if err != nil {
				panic(err)
			}

			fmt.Println("Successfully connected!")

		},
	)
	return db
}
