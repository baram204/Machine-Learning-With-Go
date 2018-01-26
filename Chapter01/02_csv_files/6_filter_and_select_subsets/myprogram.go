package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {

	// Pull in the CSV file.
	irisFile, err := os.Open("../data/iris_labeled.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// 데이터 프레임을 위한 필터를 생성한다.
	filter := dataframe.F{
		Colname:    "species", // 필터 적용될 대상 열 이름
		Comparator: "==",      //  필터 연산자
		Comparando: "Iris-versicolor", // 연산 기준값
	}

	// species 열이 "Iris-versicolor" 를 가진 레코드들만 추려서 보기 위해서
	// 필터를 적용한(거른)다.
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		log.Fatal(versicolorDF.Err)
	}

	// Output the results to standard out.
	fmt.Println(versicolorDF)

	// 한번 더 필터링 하는 데,
	// sepal_width 와 species 열만 선택한다.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})
	fmt.Println(versicolorDF)

	// 한 번 더 필터링 하는 데,
	// sepal_width 와 species 열의 첫 세 레코드(0,1,2)만 선택한다.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})
	fmt.Println(versicolorDF)

}
