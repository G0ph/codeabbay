package main

import (
	"fmt"
	"os"
)

func josephus() {

	f, _ := os.Open("text.txt")
	var n, k, count, countMain, last int

	fmt.Fscan(f, &n, &k)
	array := make([]bool, n, n)
	for {
		countMain = 0
		for i := 0; i < len(array); i++ {
			if !array[i] {
				count++
				countMain++
				last = i
				if count == k {
					array[i] = true
					//     fmt.Printf("kill %d\n",i+1)
					count = 0
				}
			}

		}

		if countMain == 1 {
			break
		}
	}

	fmt.Println(last + 1)
}
