package main

import (
	"context"
	"fmt"
	"log"

	sql "github.com/rocketlaunchr/mysql-go"
)

func singleRowQueryWithoutDbq(ctx context.Context, db *sql.DB, table string) interface{} {

	res := &store{}
	stmt := fmt.Sprintf("SELECT * FROM %s LIMIT 1", table)

	err := db.QueryRowContext(ctx, stmt).Scan(&res.ID, &res.Product, &res.Price, &res.Quantity, &res.Available, &res.Timing)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func multipleRowsQueryWithoutDbq(ctx context.Context, db *sql.DB, table string) interface{} {

	var results []interface{}
	stmt := fmt.Sprintf("SELECT * FROM %s", table)

	rows, err := db.QueryContext(ctx, stmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		res := &store{}
		err := rows.Scan(&res.ID, &res.Product, &res.Price, &res.Quantity, &res.Available, &res.Timing)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, res)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return results
}
