package main

import (
  "os"
  "fmt"
)

func calculatorMortgage(){
  f, err := os.Open("text.txt")
  if err != nil{
    fmt.Println("Error file: "+ err)
  }
  var p,r,m int
  fmt.Fscan(f,&p,&r,&m)
  fmt.Println(p,r,m)
  
}