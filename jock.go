package main

import(
  "fmt"
  "os"
  "bufio"
  "strconv"
  "strings"
)

func jock(){
  f,_ := os.Open("text.txt")
  scanner := bufio.NewScanner(f)
  scanner.Scan()
  count,_ := strconv.Atoi(scanner.Text())
  for i:=0;i<count;i++{
    scanner.Scan()
    chars := strings.Split(scanner.Text()," ")
    numbers := []int{}
    for j:=0;j<len(chars);j++{
      v,_ := strconv.Atoi(chars[j])
      numbers = append(numbers,v)
    }
    var summ int 
    for j:=0;j<len(numbers);j++{
      summ+= numbers[j]*numbers[j]
    }
    fmt.Print(summ," ")

  }

}