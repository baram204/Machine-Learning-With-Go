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
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	// 레코드(행)을 해석해서 담을 변수를 설정
	var rawCSVData [][]string

	// 레코드 하나(한 줄) 씩 읽는다.
	for {

		// READ() 로 한줄 씩 읽는다. []string 형태일 것으로 예측한다.
		record, err := reader.Read()
		// 읽고 난 뒤에는 파일의 끝(EOF) 인지 입출력에게 물어봐 체크한다.
		if err == io.EOF {
			break
		}

		// 기존 변수에, 읽어들인 새 레코드 한 줄을 추가한다.
		rawCSVData = append(rawCSVData, record)
		fmt.Println(rawCSVData)
	}

	fmt.Println(rawCSVData)
}
