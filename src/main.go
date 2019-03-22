package main

import (
	"context"
	"database/sql"
	"fmt"
	"models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test:test@tcp(185.35.187.91)/dev_emf")

	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	if db.Ping() != nil {
		fmt.Printf(err.Error())
		return
	}

	count, err := models.Locations().Count(context.Background(), db)

	fmt.Printf("Count: %d \n", count)
}
