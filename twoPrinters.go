package main

import (
	"fmt"
	"os"
)

func twoPrinters() {
	f, _ := os.Open("text.txt")
	var cases int
	fmt.Fscanln(f, &cases)
	for i := 0; i < cases; i++ {
		var x, y, n, tx, ty, t int
		fmt.Fscanln(f, &x, &y, &n)
		tx, ty = x, y
		for {
			if n <= 0 {
				break
			}
			t++
			tx--
			ty--
			if tx == 0 {
				n -= 1
				tx = x
			}
			if ty == 0 {
				n -= 1
				ty = y
			}
		}
		fmt.Print(t, " ")
	}
}
