package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/mapstructure"
	"github.com/rocketlaunchr/dbq"
)

func singleRowQueryWithDbq(ctx context.Context, db *sql.DB, table string) interface{} {
	stmt := fmt.Sprintf("SELECT * FROM %s LIMIT 1", table)

	singleRes := dbq.MustQ(ctx, db, stmt, dbq.SingleResult)
	if singleRes == nil { // no result returned
		return nil
	}

	return singleRes
}

func multipleRowsQueryWithDbq(ctx context.Context, db *sql.DB, table string, dataStruct interface{}) interface{} {
	stmt := fmt.Sprintf("SELECT * FROM %s", table)

	// Testing Multiple Data Query with dbq.MustQ
	opts := &dbq.Options{
		ConcreteStruct: dataStruct,
		DecoderConfig: &dbq.StructorConfig{
			DecodeHook:       mapstructure.StringToTimeHookFunc(time.RFC3339),
			WeaklyTypedInput: true,
		},
	}

	multResult := dbq.MustQ(ctx, db, stmt, opts)

	return multResult

}

func multipleRowsQueryWithDbqNoParseTime(ctx context.Context, db *sql.DB, table string, dataStruct interface{}) interface{} {
	stmt := fmt.Sprintf("SELECT * FROM %s", table)

	// Testing Multiple Data Query with dbq.MustQ
	opts := &dbq.Options{
		ConcreteStruct: dataStruct,
	}

	multResult := dbq.MustQ(ctx, db, stmt, opts)

	return multResult

}
