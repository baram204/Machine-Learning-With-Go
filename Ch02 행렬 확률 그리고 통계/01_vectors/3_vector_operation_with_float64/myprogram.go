package main

import (
	"fmt"
	// 벡터 연산
	"github.com/gonum/floats"
)

func main() {

	// 슬라이스로 표현되는 몇 개의 벡터를 초기화 한다.
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	// A 와 B 두 벡터의 접곱 혹은 스칼라곱을 계산한다.
	// (https://ko.wikipedia.org/wiki/스칼라곱).
	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	// A 의 각 원소에 1.5를 확대(scale, 곱)해준다
	floats.Scale(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)

	// B의 놈(norm)/길이를 계산한다.
	// https://ko.wikipedia.org/wiki/노름_공간
	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)
}
