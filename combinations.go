package main

import(
  "fmt"
  "os"
  "math/big"
)

func combi(){
  f,_ := os.Open("text.txt")
  var cases int
  fmt.Fscan(f,&cases)
  for i:=0;i<cases;i++{
      var k,n int64
      fmt.Fscan(f,&n,&k)
      del := big.NewInt(1)
      del = del.Mul(fact(k),fact(n-k))
      c := del.Div(fact(n), del)
      fmt.Print(c," ")
  }
}

func fact(f int64) (b *big.Int){
  b=big.NewInt(1)
  var j int64
  for j=1;j<=f;j++{
    a:=big.NewInt(j)
    b = b.Mul(b, a)
  }
  return b
}