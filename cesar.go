package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cesar() {
	abc := strings.Split(strings.ToUpper("abcdefghijklmnopqrstuvwxyz"), "")
	mapABC := map[int]string{}
	for i := 0; i < len(abc); i++ {
		mapABC[i] = abc[i]
	}
	f, _ := os.Open("text.txt")
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	nums := scanner.Text()
	var cases, count = strings.Split(nums, " ")[0], strings.Split(nums, " ")[1]
	h, _ := strconv.Atoi(cases)
	k, _ := strconv.Atoi(count)
	for i := 0; i < h; i++ {
		scanner.Scan()
		line := scanner.Text()
		words := strings.Split(line, " ")
		for w := 0; w < len(words); w++ {
			word := words[w]
			chars := strings.Split(word, "")
			var number int
			for j := 0; j < len(chars); j++ {
				if chars[j] == "." {
					fmt.Print(".")
					break
				}
				for k, v := range mapABC {
					if v == chars[j] {
						number = k
						break
					}
				}
				newNum := number - k
				if newNum < 0 {
					newNum = 26-newNum*(-1)
				}
				fmt.Print(mapABC[newNum])

			}
			fmt.Print(" ")
		}
	}
}
