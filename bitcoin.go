package main

import (
	"fmt"
	"os"
)

func bitCouin() {
	f, _ := os.Open("text.txt")
	var count int
	fmt.Fscanln(f, &count)
	for i := 0; i < count; i++ {
		var a, c int
		fmt.Fscan(f, &a)
		for j := 0; j < 32; j++ {
			if a&1 == 1 {
				c++
			}
			a >>= 1
		}
		fmt.Print(c, " ")
	}
}
