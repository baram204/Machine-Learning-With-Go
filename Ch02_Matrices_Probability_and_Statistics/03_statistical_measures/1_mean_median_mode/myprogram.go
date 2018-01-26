package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gonum/stat"
	"github.com/kniren/gota/dataframe"
	"github.com/montanaflynn/stats"
)

func main() {

	// Open the CSV file.
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// 데이터 프레임 새로 만든다.
	irisDF := dataframe.ReadCSV(irisFile)

	// petal_lengh 열을 가져오는데 float 타입으로 가져온다.
	sepalLength := irisDF.Col("petal_length").Float()

	// 변수의 평균(mean)을 계산한다.
	meanVal := stat.Mean(sepalLength, nil)

	// 변수의 최빈값(mode)을 계산한다.
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	// 변수의 중앙값(median)을 계산한다.
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	// 평균, 중앙값, 최빈값 필요 경우는 다음에서 보자.
	// http://blog.acronym.co.kr/401

	// 각 숫자들이 틀린 경우, 정규 분포( normal distribution )가 아닌
	// 비대칭 분포(skewed distribution) 일 가능성이..
	// http://www.statisticshowto.com/probability-and-statistics/skewed-distribution/

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median value: %0.2f\n\n", medianVal)
}
