package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
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
	fmt.Println("Db connection successfull!!")

	// Single Row Insert
	singleRowInsertWithDbq(ctx, db)
	singleRowInsertWithoutDbq(ctx, db)

	// Multiple Row Insert
	multipleRowsInsertWithDbq(ctx, db)
	multipleRowsInsertWithoutDbq(ctx, db)

	var res interface{}
	tbName := "store"

	// Single Row Query
	res = singleRowQueryWithDbq(ctx, db, tbName)
	spew.Dump(res)
	res = singleRowQueryWithoutDbq(ctx, db, tbName)
	spew.Dump(res)

	storeDataFormat := store{}

	// Multiple Rows Query
	res = multipleRowsQueryWithDbq(ctx, db, tbName, storeDataFormat)
	spew.Dump(res)
	res = multipleRowsQueryWithoutDbq(ctx, db, tbName)
	spew.Dump(res)

}
