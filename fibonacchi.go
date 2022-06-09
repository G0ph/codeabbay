package main

import (
	"fmt"
	"math/big"
	"os"
)

func fibonacchi() {
	f, _ := os.Open("text.txt")
	var count int
	fmt.Fscan(f, &count)
	for i := 0; i < count; i++ {
		var a big.Int
		fmt.Fscan(f, &a)
		fmt.Print(func(chis big.Int) (index int) {
			if chis.String() == "0" {
				return 0
			}
			if chis.String() == "1" {
				return 1
			}
			index = 2
			var a, b = big.NewInt(0), big.NewInt(1)
			for {
				a = a.Add(a, b)
				if a.Cmp(&chis) == 0 {
					return
				}
				index++
				a, b = b, a
			}
		}(a), " ")

	}
}

//942 219 776 745 38 509 959 367 358 920 733 592 142 533 666 906 882 419 35
