// Script that is used to populate mysql database with diffrnt character up to a thousand combination

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Pallinder/go-randomdata"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rocketlaunchr/dbq"
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

	var stmt string
	// Insert Batch data of 100000 in total hops of 50 at 2000 per hop
	totalDataNum := 100000 // Total number of data that is to be sent to db
	maxBatchPerHop := 2000 // total number of row to insert per hop

	hCount := 0
	total := 0
	dataValues := make([]string, maxBatchPerHop)
	index := 0
	for i := 0; i < totalDataNum; i++ {

		fullName := randomdata.FullName(randomdata.RandomGender)
		dataValues[index] = fullName

		if index == maxBatchPerHop-1 { // This is for a batch insert of 2000 values per hop
			hCount++
			total += maxBatchPerHop

			stmt = dbq.INSERT("benchmark", []string{"name"}, len(dataValues), dbq.MySQL)
			dbq.MustE(ctx, db, stmt, nil, dataValues)

			fmt.Printf("\n%d data inserted in hop %d\nTotal row Inserted: %d\n",
				len(dataValues), hCount, total)
			// reset data container and index pointer
			dataValues = make([]string, maxBatchPerHop)
			index = 0
			continue
		}

		index++
	}

}
