package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/rocketlaunchr/dbq"
)

func singleRowInsertWithDbq(ctx context.Context, db *sql.DB) {
	// Single Row Insert Data with dbq.E
	newStore := []interface{}{
		[]interface{}{"404", "motor Cycle", 36, 4464646.46, 1, time.Now()},
	}

	stmt := dbq.INSERT("store", []string{"Id", "product", "quantity", "price", "available", "timing"}, len(newStore), dbq.MySQL)

	dbq.MustE(ctx, db, stmt, nil, newStore)
}

func multipleRowsInsertWithDbq(ctx context.Context, db *sql.DB) {
	// Multiple Row Insert with dbq.E and dbq.INSERT
	newStores := []interface{}{
		[]interface{}{"405", "Comic Book", 25, 456.34, 1, time.Now()},
		[]interface{}{"406", "Movie Ticket", 17, 1250.5, 1, time.Now()},
		[]interface{}{"407", "Teddy Bear", 30, 99.99, 0, time.Now()},
	}

	stmt := dbq.INSERT("store", []string{"Id", "product", "quantity", "price", "available", "timing"}, len(newStores), dbq.MySQL)

	dbq.MustE(ctx, db, stmt, nil, newStores)
}
