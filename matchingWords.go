package main

import (
	"fmt"
	"os"
  "bufio"
  "strings"
  "sort"
)

func matchWords() {
	f, _ := os.Open("text.txt")
  scanner := bufio.NewScanner(f) 
  scanner.Scan()
  words := strings.Split(scanner.Text()," ")
  card := map[string]int{}
  for i:=0;i<len(words);i++{
    card[words[i]]++
  }
  wordsOnes := []string{}
  for k,v := range card{
      if v>1{
        wordsOnes = append(wordsOnes, k)
      }
    }
  sort.Strings(wordsOnes)  
  
  for i:=0;i<len(wordsOnes);i++{
    fmt.Print(wordsOnes[i]," ")
  }
}