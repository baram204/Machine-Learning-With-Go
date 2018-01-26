package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris_unexpected_fields.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// 한 행에 5 개의 필드가 있어야 한다고 설정한다.
	// 즉 5개가 있는지 유효성 검증을 할 수 있게 되는 것이다.
	reader.FieldsPerRecord = 5

	// rawCSVData will hold our successfully parsed rows.
	var rawCSVData [][]string

	// Read in the records one by one.
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 해석 오류가 있으면 출력하고 계속 진행한다.
		if err != nil {
			log.Println(err)
			continue
		}

		// Append the record to our data set, if it has the expected
		// number of fields.
		rawCSVData = append(rawCSVData, record)
	}

	fmt.Printf("parsed %d lines successfully\n", len(rawCSVData))
}
