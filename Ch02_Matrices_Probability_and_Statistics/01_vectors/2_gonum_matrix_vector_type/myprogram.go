package main

import (
	"fmt"
	// 고넘이 제공하는 float64 타입을 사용하는 행렬 타입
	"github.com/gonum/matrix/mat64"
)

func main() {

	// 새로운 벡터 값을 생성한다.
	myvector := mat64.NewVector(2, []float64{11.0, 5.2})

	fmt.Println(myvector)
}
