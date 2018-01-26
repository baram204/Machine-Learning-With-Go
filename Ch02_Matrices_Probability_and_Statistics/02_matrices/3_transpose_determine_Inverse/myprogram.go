package main

import (
	"fmt"
	"log"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create a new matrix a.
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// 행렬을 전치(transpose of matrix)한 결과를 계산한다.
	// https://ko.wikipedia.org/wiki/전치행렬
	ft := mat64.Formatted(a.T(), mat64.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// 행렬식을 계산한다.
	deta := mat64.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// a 역행렬을 구하고 출력한다.
	// 행렬 A의 역행렬은 A와 곱해서 항등행렬 E가 나오는 행렬을 A의 역행렬
	aInverse := mat64.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat64.Formatted(aInverse, mat64.Prefix("       "))
	fmt.Printf("a^-1 = %v\n\n", fi)
}
