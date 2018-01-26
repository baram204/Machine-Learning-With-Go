package main

import "fmt"

func main() {

	// 벡터를 슬라이스를 통해 초기화한다.
	var myvector []float64

	// 몇 몇 원소를 벡터에 추가한다.
	myvector = append(myvector, 11.0)
	myvector = append(myvector, 5.2)

	fmt.Println(myvector)
}
