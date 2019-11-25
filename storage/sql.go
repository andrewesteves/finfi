package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	user = "root"
	pass = "4321"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/finfi?parseTime=true", user, pass))
	if err != nil {
		panic(err.Error())
	}
	return db
}
