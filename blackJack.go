package main
import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "strconv"
)
const(
  
)


func blackJack(){
  
  f,_ := os.Open("text.txt")
  scanner := bufio.NewScanner(f)
  scanner.Scan()
  for scanner.Scan() {
    line := strings.Split(scanner.Text()," ")
    fmt.Print(getSumm(line)," ")
  }
}

func getSumm(array []string) (string){
  var summ int 
  for i:=0;i<len(array);i++{
    switch array[i]{
      case "T","J","K","Q":
        summ+=10
      case "A":
        if summ+11 > 21{
          summ+=1
        } else{
          summ+=11
        }
        default:
          chis,_:= strconv.Atoi(array[i])
          summ+=chis
    }
  }
  if summ > 21{
    return "bust"
  }
  return fmt.Sprint(summ)
}