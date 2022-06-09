package main

import (
	"fmt"
	"os"
	"strconv"
)

type Figure struct {
	ch  int
	num int
}

func KandQ() {
	f, _ := os.Open("text.txt")
	var cases int
	fmt.Fscan(f, &cases)
	array := [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for i := 0; i < cases; i++ {
		var ki, qu string
		fmt.Fscan(f, &ki, &qu)
		var kChar, qChar, kNum, qNum int

		for i := 0; i < 8; i++ {
			if array[i] == string(qu[0]) {
				qChar = i
			}
			if array[i] == string(ki[0]) {
				kChar = i
			}
		}
		kNum, _ = strconv.Atoi(string(ki[1]))
		qNum, _ = strconv.Atoi(string(qu[1]))

		k := Figure{
			ch:  kChar,
			num: kNum,
		}

		q := Figure{
			ch:  qChar,
			num: qNum,
		}

		if k.ch == q.ch || k.num == q.num {
			fmt.Print("Y ")
			continue
		}
		switch {
		case q.ch > k.ch && q.num > k.num:
			raz := q.num - k.num
			if q.ch-raz == k.ch {
				fmt.Print("Y ")
				continue
			}

		case q.ch > k.ch && q.num < k.num:
			raz := k.num - q.num
			if q.ch-raz == k.ch {
				fmt.Print("Y ")
				continue
			}

		case q.ch < k.ch && q.num > k.num:
			raz := q.num - k.num
			if k.ch-raz == q.ch {
				fmt.Print("Y ")
				continue
			}

		case q.ch < k.ch && q.num < k.num:
			raz := k.num - q.num
			if k.ch-raz == q.ch {
				fmt.Print("Y ")
				continue
			}

		}
		fmt.Print("N ")
	}
}
