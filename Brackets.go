package main

import "fmt"
import "os"
import "regexp"
import "bufio"
import "strings"

const S = ""

func brackets2() {
	f, _ := os.Open("text.txt")
	var count int
	fmt.Fscan(f, &count)
	scanner := bufio.NewScanner(f)
	for q := 0; q < count; q++ {
		scanner.Scan()
		str := scanner.Text()
		var re = regexp.MustCompile(`[^\[\]{}()<>]`)
		s := re.ReplaceAllString(str, "")
		s1 := strings.Split(s, "")

		op := []string{"(", "{", "[", "<"}
		cl := []string{")", "}", "]", ">"}
		correct := 1
	Loop:
		for i := 0; i < len(s1); i++ {
			for j := 0; j < 4; j++ {
				if cl[j] == s1[i] {
					if i-1 < 0 {
						correct = 0
						break Loop
					}
					if s1[i-1] == op[j] {
						s1[i], s1[i-1] = S, S
					} else {
						for k := i - 1; k >= 0; k-- {
							if s1[k] == S {
								continue
							}
							if s1[k] != op[j] {
								correct = 0
								break Loop
							} else {
								s1[k], s1[i] = S, S
								break
							}
						}
					}
				}
			}
		}
		for i := 0; i < len(s1); i++ {
			if s1[i] != S {
				correct = 0
			}
		}
		fmt.Print(correct, " ")

	}
}

// func tBracks(str []string,n int, pred,obr string, in map[string]int, other []string) (correct byte) {
//   fmt.Println(pred,obr,in,other)
//   correct=1
//   switch string(pred){
//             case string(other[0]),string(other[1]),string(other[2]):
//               fmt.Println("I'm here",pred)
//               correct = 0
//               return
//             default:
//               in[obr]--
//               if in[obr]<0{
//                 correct = 0
//                 fmt.Printf("%s: %d\n",obr,in[obr])
//                 return
//               }
//               for i:=n-1;i<=0;i--{

//               }

//   }
//   return
// }

//{[]<()>{({})[{()]}[<>[][]{({})}{}<>]()}}
//{{[{]}[<>[][]{({})}{}<>]()}}
