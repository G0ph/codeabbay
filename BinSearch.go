package main

import (
  "fmt"
  "os"
)

func binSearch(){
  f,_ := os.Open("text.txt")  
  var cases int
  fmt.Fscan(f,&cases)
  for i:=0;i<cases;i++{
    var a,b,c,d float64
    fmt.Fscan(f, &a,&b,&c,&d)
    goal := a* //A * x + B * sqrt(x ^ 3) - C * exp(-x / 50) - D = 0
    
  }
}