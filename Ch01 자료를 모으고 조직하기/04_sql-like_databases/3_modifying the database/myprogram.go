package main

import (
	"database/sql"
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

	// Open a database value. Specify the postgres driver
	// for databases/sql.
	db, err := sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 어떤 값들을 갱신한다.
	// 이것은 변경(modification)이기 때문에 Exec 를 사용해서 변경하는 쿼리문을 날린다.
	res, err := db.Exec("UPDATE iris SET species = 'setosa' WHERE species = 'Iris-setosa'")
	if err != nil {
		log.Fatal(err)
	}

	// 얼마나 많은 행이 영향을 받았는지 본다.
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	// Output the number of rows to standard out.
	log.Printf("affected = %d\n", rowCount)
}
