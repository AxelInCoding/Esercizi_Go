/*
Leggere da standard input una sequenza di "parole",  
fino a incontrare la parola "stop", e stampare 
le parole del testo ("stop" esclusa), separate 
da spazi, senza punteggiatura ( : , . ; ? ! ). 
Ad esempio, il testo:
ciao, come stai? stop
andrà stampato così: 
ciao come stai
*/

package main 
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

//pulisce una stringa da carattere speciali
func pulisci( w string ) string {
  var builder strings.Builder // per evitare di utilizzare la concatenazione di stringhe
  for _, r := range w {
    if unicode.IsPunct(r){
      continue
    }
    builder.WriteRune(r) //scrivo la runa che NON è un carattere di punteggiatura nel Builder
  }
  return builder.String() //converto contenuto in una stringa
}

func main(){
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  first := true

  for scanner.Scan() {
    w := scanner.Text()
    if w == "stop" {
      break
    }
    p := pulisci(w)
    if p != "" {
      if !first {
        fmt.Print(" ")
      }
      fmt.Print(p)
      first = false
    }
  }
}
