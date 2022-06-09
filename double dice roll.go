package main

import (
	"fmt"
	"os"
)

func doubleDice() {
	f, _ := os.Open("text.txt")
	var cases int
	fmt.Fscan(f, &cases)
	for i := 0; i < cases; i++ {
		var a, b int
		fmt.Fscan(f, &a, &b)
		fmt.Print((a%6+1)+(b%6+1), " ")
	}

}
