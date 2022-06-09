package main
import(
  "fmt"
  
)
func selectSort(){
  var cases int
  fmt.Scan(&cases)
  var el int
  slice := []int{}
  indexes := []int{}
  for i:=0;i<cases;i++{
    fmt.Scan(&el)
    slice = append(slice,el)
  }
  for {
    max := slice[0]
    ind := 0
    for j:=1;j<cases;j++{
      if slice[j] > max {
        max = slice[j]
        ind = j
      }
    }
    
    slice[cases-1],slice[ind] = slice[ind],slice[cases-1]
    cases--
    
    indexes =append(indexes,ind)
    if cases == 1{
      break
    }
  }
  fmt.Println(indexes)
}