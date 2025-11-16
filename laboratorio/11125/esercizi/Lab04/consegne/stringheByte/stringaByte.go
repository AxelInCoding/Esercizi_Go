package main 
import "fmt"

func main (){
  var s string;
  fmt.Print("Inserisci una stringa: ")
  fmt.Scan(&s)

  if len(s) < 2 {
    fmt.Print("La stringa deve avere almeno due caratteri.")
    return
  }

  for i:=1; i < len(s); i++ {
    if s[i] > s [i-1] {
      fmt.Print("+")
    } else if s[i] < s[i-1] {
      fmt.Print("-")
    } else {
      fmt.Print("=")
    }
  }

}
