package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// pq is the libary that allows us to connect
	// to postgres with databases/sql.
	_ "github.com/lib/pq"
)

func main() {

	// Get my postgres connection URL. I have it stored in
	// an environmental variable.
	pgURL := os.Getenv("PGURL")
	if pgURL == "" {
		log.Fatal("PGURL empty")
	}

	// Open a database value.  Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/*
		1. Query : selects, groups or aggregate data and returns rows of data to us
		2. Exec : update, inserts or otherwise modifies the state of the database
				without expectation tha portions of the data stored in th database should be returned
				Exec 는 데이터를 수정하는데, 수정한 결과를 반환하지 않는다.
	 */


	// Query the database.
	rows, err := db.Query(`
		SELECT 
			sepal_length as sLength, 
			sepal_width as sWidth, 
			petal_length as pLength, 
			petal_width as pWidth 
		FROM iris
		WHERE species = $1`, "Iris-setosa")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate over the rows, sending the results to
	// standard out.
	// 쿼리셋의 Next () 하면 0 번 부터 하나씩 포인터를 하나씩 증가시킨다.
	for rows.Next() {

		var (
			sLength float64
			sWidth  float64
			pLength float64
			pWidth  float64
		)

		// scan 을 통해서 쿼리셋의 한 레코드를 특정 변수에 넣는다.
		if err := rows.Scan(&sLength, &sWidth, &pLength, &pWidth); err != nil {
			log.Fatal(err)
		}

		// 값이 입력된 변수들을 출력한다.
		fmt.Printf("%.2f, %.2f, %.2f, %.2f\n", sLength, sWidth, pLength, pWidth)
	}

	// Check for errors after we are done iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
