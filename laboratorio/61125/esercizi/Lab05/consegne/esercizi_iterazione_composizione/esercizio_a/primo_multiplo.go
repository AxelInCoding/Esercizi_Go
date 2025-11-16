package main 
import "fmt"


func main (){
  var n int
  found := false

  for i:= 0; i < 5; i++ {
    fmt.Scan(&n)

    if n < 0 {
      fmt.Println("Numero non valido, deve essere maggiore di 0.")
    }

    if n % 3 == 0 {
      fmt.Println(n)
      found = true;
      break
    }

  }

  if !found {
    fmt.Println("no")
  }

}
