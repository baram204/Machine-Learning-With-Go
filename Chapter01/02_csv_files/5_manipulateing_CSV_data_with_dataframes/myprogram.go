package main

import (
	"fmt"
	"log"
	"os"
	// standarized way to filter, subset, select portions of tabular datasets.\
	// > go get github.com/kniren/gota/dataframe
	"github.com/kniren/gota/dataframe"
)

func main() {

	// Open the CSV file.
	irisFile, err := os.Open("../data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// CSV 파일로부터 데이터프레임을 생성한다.
	// 열의 타입은 유추될(inferred) 것이다.
	irisDF := dataframe.ReadCSV(irisFile)

	// 건전성 확인 후, stdout 으로 레코드들을 보여준다.
	// Gota 가 데이터프레임을 프리티 출력해줄 것이다.
	fmt.Println(irisDF)
}
