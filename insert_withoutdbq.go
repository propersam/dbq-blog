package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func singleRowInsertWithoutDbq(ctx context.Context, db *sql.DB) {
	// Without DBQ
	newData := []interface{}{"400", "motor Cycle", 36, 446.46, 1, time.Now()}
	insertQuery := "INSERT INTO store(id, product, quantity, price, available, timing) VALUES (?, ?, ?, ?, ?, ?)"

	_, err := db.ExecContext(ctx, insertQuery, newData...)
	if err != nil {
		log.Fatal(err)
	}

}

func multipleRowsInsertWithoutDbq(ctx context.Context, db *sql.DB) {
	// multiple data to insert
	newStores := [][]interface{}{
		[]interface{}{"401", "Comic Book", "25", "456.34", "1", time.Now()},
		[]interface{}{"402", "Movie Ticket", "17", "1250.5", "1", time.Now()},
		[]interface{}{"403", "Teddy Bear", "30", "99.99", "0", time.Now()},
	}

	insertQuery := "INSERT INTO store(id, product, quantity, price, available, timing) VALUES (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?)"

	spreadData := []interface{}{}

	for _, data := range newStores {
		for _, col := range data {
			spreadData = append(spreadData, col)
		}
		
	}

	_, err := db.ExecContext(ctx, insertQuery, spreadData...)
	if err != nil {
		log.Fatal(err)
	}
}
