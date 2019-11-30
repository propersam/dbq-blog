package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/rocketlaunchr/dbq"
)

func BenchmarkSingleRowQueryWithoutDBQ(b *testing.B) {
	ctx := context.Background()

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "db-key", "127.0.0.1", 3306, "dataframe")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	table := "benchmark"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := &bnchmark{}
		stmt := fmt.Sprintf("SELECT * FROM %s LIMIT 1", table)

		err := db.QueryRowContext(ctx, stmt).Scan(&res.ID, &res.Name, &res.Timestamp)
		if err != nil {
			log.Fatal(err)
		}

		_ = res
	}

}

func BenchmarkMultipleRowsQueryWithoutDBQ(b *testing.B) {
	ctx := context.Background()

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "db-key", "127.0.0.1", 3306, "dataframe")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	table := "benchmark"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var results []interface{}
		stmt := fmt.Sprintf("SELECT * FROM %s", table)

		rows, err := db.QueryContext(ctx, stmt)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			res := &bnchmark{}
			err := rows.Scan(&res.ID, &res.Name, &res.Timestamp)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, res)
		}

		err = rows.Err()
		if err != nil {
			log.Fatal(err)
		}

		_ = results
	}

}

func BenchmarkSingleRowQueryWithDBQ(b *testing.B) {
	ctx := context.Background()

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "db-key", "127.0.0.1", 3306, "dataframe")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	table := "benchmark"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stmt := fmt.Sprintf("SELECT * FROM %s LIMIT 1", table)

		singleRes := dbq.MustQ(ctx, db, stmt, dbq.SingleResult)

		_ = singleRes
	}

}

func BenchmarkMultipleRowsQueryWithDBQ(b *testing.B) {
	ctx := context.Background()

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "db-key", "127.0.0.1", 3306, "dataframe")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	table := "benchmark"
	benchDataFormat := bnchmark{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := multipleRowsQueryWithDbq(ctx, db, table, benchDataFormat)
		_ = res
	}

}

func BenchmarkMultipleRowsQueryWithDBQNoTimeParse(b *testing.B) {
	ctx := context.Background()

	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", "root", "db-key", "127.0.0.1", 3306, "dataframe")
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	table := "benchmark"
	benchDataFormat := bnchmark2{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := multipleRowsQueryWithDbqNoParseTime(ctx, db, table, benchDataFormat)
		_ = res
	}

}
