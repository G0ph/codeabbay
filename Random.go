package main

//A, C, M, X0, N
import (
	"fmt"
	"os"
)

func myRand() {
	f, _ := os.Open("text.txt")
	var count int
	fmt.Fscan(f, &count)
	for j := 0; j < count; j++ {
		var a, c, m, x, n int
		fmt.Fscan(f, &a, &c, &m, &x, &n)
		//Xnext = (A * Xcur + C) % M
		for i := 0; i < n; i++ {
			x = (a*x + c) % m
		}
		fmt.Print(x, " ")
	}
}
