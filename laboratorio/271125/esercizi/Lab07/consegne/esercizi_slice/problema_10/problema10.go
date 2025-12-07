/*
Leggere da standard input una sequenza di "parole",
 fino a incontrare EOF, e stampare solo (e tutte)
  le parole che hanno lunghezza massima, 
  una parola per riga.
*/

package main 
import (
    "bufio"
    "fmt"
    "os"
)

func main(){
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)

  maxLen := 0
  maxWords := []string{}
  
  for scanner.Scan() {
    w := scanner.Text()
    l := len([]rune(w)) //lunghezza di una singola parola
    if l > maxLen {
      maxLen = l
      maxWords = []string{w}
    } else if l == maxLen {
      maxWords = append(maxWords, w)
    }
  }

  for _, w := range maxWords {
    fmt.Println(w)
  }
}
