package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// 우리 행렬을 위한 평평한 자료를 만든다.
	data := []float64{1.2, -5.7, -2.4, 7.3}

	// 밀집행렬 형태의 새행렬을 2x2 로 만든다.
	// 밀집행렬(dense matrix)는 희소행렬(sparse matrix) 의 반대다.
	// 희소는 대부분의 원소가 0인 경우를 말한다.
	a := mat64.NewDense(2, 2, data)

	// As a sanity check, output the matrix to standard out.
	fa := mat64.Formatted(a, mat64.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}
