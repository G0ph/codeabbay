package main

import (
	"fmt"
	"os"
)

func getWeather() {
	f, _ := os.Open("text.txt")
	var c int
	fmt.Fscan(f, &c)
	array := make([]float64, c)
	array2 := make([]float64, c)
	for i := 0; i < c; i++ {
		var a float64
		fmt.Fscan(f, &a)
		array[i] = a
	}
	array2[0], array2[len(array)-1] = array[0], array[len(array)-1]
	for i := 1; i < c-1; i++ {
		array2[i] = (array[i-1] + array[i] + array[i+1]) / 3
	}
	for i := 0; i < c; i++ {
		fmt.Printf("%.12g ", array2[i])
	}

}
