package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// CSVRecord 타입은 성공적으로 해석됐을 때의 행을 포함한다.
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

func main() {

	// Open the iris dataset file.
	f, err := os.Open("../data/iris_mixed_types.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// 구조체를 한 행으로하는, 여러 행을 가질 수 있는 자료를 선언한다.
	var csvData []CSVRecord

	// 라인은 로깅을 추적하는데 도움을 준다.
	line := 1

	// Read in the records looking for unexpected types.
	for {

		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// 하나의 행을 위한 변수를 선언
		var csvRecord CSVRecord

		// 기대되는 타입을 근거로 해서, 한 행(레크도)씩 각 값을 해석한다.
		for idx, value := range record {

			// 레코드의 4번째 필드를 읽을 때는 문자열로 변환한다.
			if idx == 4 {

				// 빈 문자열인지 확인해서 비어있으면 해석 루프를 중단한다.
				if value == "" {
					log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
					csvRecord.ParseError = fmt.Errorf("Empty string value")
					break
				}

				// Add the string value to the CSVRecord.
				csvRecord.Species = value
				continue
			}

			// 인덱스가 4가 아닌, 다른 것들(otherwise)에 대해서는 float64 타입으로 값을 해석한다.
			var floatValue float64

			// 만약 값이 플롯으로 해석이 안 되면, 기록을 남기고 해석 루프를 중단한다.
			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				log.Printf("Parsing line %d failed, unexpected type in column %d\n", line, idx)
				csvRecord.ParseError = fmt.Errorf("Could not parse float")
				break
			}

			// CVSRecord 변수에다가 각각의 해석 값을 저장한다.
			// 인덱스 번호에 맞춰서.
			switch idx {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.PetalLength = floatValue
			case 3:
				csvRecord.PetalWidth = floatValue
			}
		}

		// 성공적으로 해석된 레코드 하나를 전체 레코드를 담으려고 생성한 슬라이스에 적용한다.
		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}

		// 줄 숫자세는 카운터를 증가시킨다.
		line++
	}

	// 결과를 출력한다.
	fmt.Printf("successfully parsed %d lines\n", len(csvData))
}
