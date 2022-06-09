package main
import(
  "os"
  "fmt"
  "math"
  "strings"
)
func quadratic(){
  f,_ := os.Open("text.txt")
  var cases int
  fmt.Fscan(f,&cases)
  var s,sep string
  for i:=0;i<cases;i++{
    var a,b,c float64
    fmt.Fscan(f,&a,&b,&c)
    d:= b*b-4*a*c
    if d>=0{
      x1 := (-1*b + math.Sqrt(d)) / (2*a)
      x2 := (-1*b - math.Sqrt(d)) / (2*a)
      s += sep+fmt.Sprint(x1,x2)
    }else{
      x1 := (-1*b)/(2*a)
      x2 := (-1*b)/(2*a)
      imag := math.Sqrt(-1*d)/(2*a)
      s += sep+fmt.Sprint(complex(x1,imag),complex(x2,imag*-1))
    }
        
    sep = "; "
  }
  s = strings.ReplaceAll(s,"(","")
  s = strings.ReplaceAll(s,")","")
  fmt.Print(s)
  
}