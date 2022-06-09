package main

// b1DRfcpyIsL1Ole9hMsZwvH9 TUqZb2K  Rq3Zkv l xTQSDkEsaPW0 41gSpJ6EYFaS.
// b1DRfcpyIsL1Ole9hMsZwvH9 TUqZb2K  Rq3Zkv l xTQSDkEsaPW0 41gSpJ6EYFaS.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partControl() {
	f, _ := os.Open("text.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	words := strings.Split(line, " ")

	for i := 0; i < len(words); i++ {
		symb, _ := strconv.Atoi(words[i])
		var count byte
		var tmp = symb
		for j := 0; j < 8; j++ {
			if tmp&1 == 1 {
				count++
			}
			tmp >>= 1
		}
		if count%2 != 0 {
			continue
		}
		if symb&128 == 128 {
			symb = symb ^ 128
		}
		fmt.Print(string(symb))

	}
}
