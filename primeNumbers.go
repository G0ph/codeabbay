package main

import (
	"fmt"
	"math/big"
	"os"
)

const countels = 9999999

func primeNumbers() {
	var array [countels]int64
	getPrimeArray(&array)
	f, _ := os.Open("text.txt")
	var cases int
	fmt.Fscanln(f, &cases)

	for i := 0; i < cases; i++ {
		var num int
		fmt.Fscan(f, &num)
		c := 1
		for i := 1; i < countels; i++ {
			if array[i] != 0 {
				if c == num {
					fmt.Printf("%d ", array[i])
				}
				c++
			}
		}
	}
}

//10
//105637 194098 134067 197015 150437 121520 174558 129853 134232 184278

func getPrimeArray(array *[countels]int64) {
	createArray(array)
	var j int
	for j = 1; j < countels; j++ {
		if !big.NewInt(array[j]).ProbablyPrime(0) {
			array[j] = 0
		}
	}
}

func createArray(array *[countels]int64) {
	var i int64
	for i = 0; i < countels; i++ {
		array[i] = i
	}
}
