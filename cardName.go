package main

import(
  "os"
  "fmt"
)

func cardNames(){
  suits := []string{"Clubs", "Spades", "Diamonds", "Hearts"}
  ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
  f,_ := os.Open("text.txt")
  var cases int
  fmt.Fscan(f,&cases)
  for i:=0;i<cases;i++{
    var card, suit, rang byte
    fmt.Fscan(f,&card)
    suit = card/13
    rang = card%13
    fmt.Printf("%s-of-%s ",ranks[rang],suits[suit])
  }
}