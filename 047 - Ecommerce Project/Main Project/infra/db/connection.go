package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString() string {
	// User -> postgres
	// Password -> admin
	// host -> localhost
	// Port -> 5432
	// DB_Name -> go_ecommerce

	return "user=postgres password=admin host=localhost port=5432 dbname=go_ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString()

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return dbCon, nil
}
