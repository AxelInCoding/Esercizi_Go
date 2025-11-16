package main 
import "fmt"
import "strings"


func horizLine(ch byte, n int) string {
  return " " + strings.Repeat(string(ch), n) + " "
}

func main() {
  var parola, maxParola string

  fmt.Print("Inserisci delle parole (scrivi 'stop' per terminare): ")
  
  for {
    fmt.Scan(&parola)

      if parola == "stop" {
        break
      }

      if len(parola) > len(maxParola){
        maxParola = parola
      }
  }

  largghezza := len (maxParola) + 2

  fmt.Println(horizLine('-', largghezza))
  fmt.Printf("| %s |\n", maxParola)
  fmt.Println(horizLine('-', largghezza))
}
