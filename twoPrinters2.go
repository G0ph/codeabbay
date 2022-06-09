package main

import (
	"fmt"
	"os"
)

func twoPrinters2() {
	f, _ := os.Open("text1.txt")
	var cases int
	fmt.Fscanln(f, &cases)
	for i := 0; i < cases; i++ {
		var x, y, n, timeAll int
		time := make(chan int, 2)
		pages := make(chan int, 2)

		fmt.Fscanln(f, &x, &y, &n)
		go printPage(x, pages, time)
		go printPage(y, pages, time)

		for j := 0; j < n; j++ {
			pages <- 1
			timeAll += <-time
		}

		fmt.Print(timeAll, "")
	}
}

func printPage(t int, p, time chan int) {
	<-p
	time <- t
}
