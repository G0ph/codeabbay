package main

import(
  "fmt"
  "os"
  "strings"
)

func bAndC(){
  f,_ := os.Open("text.txt")
  var number string
  var count int
  fmt.Fscan(f,&number,&count)
  numberArray := strings.Split(number, "")
  for i:=0;i<count;i++{
    var chis string
    fmt.Fscan(f,&chis)
    chisArray := strings.Split(chis, "")
    var posCount, chisCount int
    for j:=0;j<4;j++{
      for k:=0;k<4;k++{
        if chisArray[j] == numberArray[k]{
          if j==k{
            posCount++
          }else{
            chisCount++
          }
        }

      }
    }
    fmt.Print(posCount,"-",chisCount," ")

    
  }
}