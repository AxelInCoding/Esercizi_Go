package main 
import(
  "fmt"
  "bufio"
  "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("un numero: ")

    scanner.Scan()
    n := scanner.Text()

    dict := make(map[rune]string)

    for _, digit := range n {
    // se la numero della crifra non è stato finora incontrato
     if _, ok := dict[digit]; !ok { 
      fmt.Printf("parola per %c ?", digit)
      scanner.Scan()
      word := scanner.Text()
      dict[digit] = word
     } 
    }

    //Stampa sequenza delle parole
    for i, digit := range n {
      fmt.Print(dict[digit])
      if i < len(n) - 1 {
        fmt.Print(" - ")
      }
    }
    fmt.Println()
}
