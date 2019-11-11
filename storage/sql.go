package storage

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:4321@tcp(127.0.0.1:3306)/finfi?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
