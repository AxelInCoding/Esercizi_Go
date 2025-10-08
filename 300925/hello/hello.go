package main
import "fmt"

func main (){
  var n int
  fmt.Print("Quante volte vuoi esser salutato padrone?")
  fmt.Scan(&n)

  for i:=0; i<n; i++ {
  fmt.Println("Saluto ", n)
  }
}
