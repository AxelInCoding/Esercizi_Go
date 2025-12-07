/*
Leggere da standard input una sequenza di "parole",
 fino a incontrare la parola "stop", e stampare le
 parole in ordine lessicografico ("stop" esclusa), 
 una parola per riga.
*/

package main
import (
  "fmt"
  "bufio"
  "os"
  "sort"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  words := []string{}

  for scanner.Scan() {
    w := scanner.Text()
    if w == "stop" {
      break
    }
    words = append(words, w)
  }

  sort.Strings(words) //ordina in modo lessicografico

  for _, w := range words {
    fmt.Println(w)
  }
}
