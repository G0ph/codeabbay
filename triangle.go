package main

import (
	"fmt"
	"math"
	"os"
)

func triangle() {
	f, _ := os.Open("text.txt")
	var cases int
	fmt.Fscan(f, &cases)
	for i := 0; i < cases; i++ {
		var X1, Y1, X2, Y2, X3, Y3 float64
		fmt.Fscan(f, &X1, &Y1, &X2, &Y2, &X3, &Y3)
		a := math.Sqrt((X1-X2)*(X1-X2) + (Y1-Y2)*(Y1-Y2))
		b := math.Sqrt((X2-X3)*(X2-X3) + (Y2-Y3)*(Y2-Y3))
		c := math.Sqrt((X1-X3)*(X1-X3) + (Y1-Y3)*(Y1-Y3))
		s := (a + b + c) / 2
		A := math.Sqrt(s * (s - a) * (s - b) * (s - c))
		fmt.Printf("%.100g ", A)
	}

}
