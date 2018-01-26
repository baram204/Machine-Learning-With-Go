package main

import (
	"fmt"
	//"log"

	"github.com/gonum/matrix/mat64"
	"math"
)

func main() {

	// Create a new matrix a.
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat64.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})

	// 크기가 다른 행렬을 만든다.
	c := mat64.NewDense(3, 2, []float64{3, 2, 1, 4, 0, 8})

	// 빈 행렬을 만들고 a,b 행렬을 더한다.
	d := mat64.NewDense(0, 0, nil)
	d.Add(a, b)
	fd := mat64.Formatted(d, mat64.Prefix("      "))
	fmt.Printf("d=a+b=%0.4v\n\n", fd)

	//곱한다.
	f := mat64.NewDense(0, 0, nil)
	f.Mul(a, c)

	ff := mat64.Formatted(f, mat64.Prefix("     "))
	fmt.Printf("f=ac=%0.4v\n\n", ff)

	// 거듭제곱(power)
	g := mat64.NewDense(0, 0, nil)
	g.Pow(a, 5)
	fg := mat64.Formatted(g, mat64.Prefix("      "))
	fmt.Printf("g=a^5=%0.4v\n\n", fg)

	// a 원소에 각각 함수를 적용한다.
	h := mat64.NewDense(0, 0, nil)
	sqrt := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
	h.Apply(sqrt, a)
	fh := mat64.Formatted(h, mat64.Prefix("          "))
	fmt.Printf("h=sqrt(a)=%0.4v\n\n", fh)
}
