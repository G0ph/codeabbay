package main

import (
	"fmt"
	"math"
	"os"
)

func bankAcc() {
	f, err := os.Open("text.txt")
	if err != nil {
		panic(err)
	}
	var start, goal, p float64
	var c int
	fmt.Fscanln(f, &c)
	for i := 0; i < c; i++ {
		fmt.Fscan(f, &start, &goal, &p)

		var y int
		for start <= goal {
			y++
			start += start * p / 100
			start = math.Floor(start*100) / 100

		}
		fmt.Print(y, " ")
	}
}

//137 163 273 302 12 137 57 324 273 164 273 109 302 163 27 102 163
