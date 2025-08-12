package config

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", "root:12345@tcp(localhost:3306)/orfi?parseTime=true")

		if err != nil {
			log.Fatal("Error conectando a la base de datos:", err)
		}
	})
	return db
}
