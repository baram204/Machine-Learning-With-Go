package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	// 데이터 셋 파일 열면서 결과와 오류 받기.
	f, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	// 모든 작업이 마친 후 열린 파일을 닫히도록 미리 작업을 지연시켜 놓는다.
	defer f.Close()

	// 열린 파일을 가지고 새 리더를 생성한다.
	reader := csv.NewReader(f)


	// 레코드(한 행)에 몇 개의 필드(열)이 있는 지 모른다고 가정한다.
	// -1 을 지정하면 필드의 개수를 지정하지 않는다.
	reader.FieldsPerRecord = -1

	// 모두 읽는다.
	// 결과의 타입은 [][]string 이다.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawCSVData)
}
