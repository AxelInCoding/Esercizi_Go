/*
Leggere da standard input una sequenza di 
"parole",  fino a incontrare la parola "stop",
 e stampare prima le parole
 di lunghezza pari ("stop" esclusa), su 
 una riga, separate da spazi, e poi le parole 
 di lunghezza dispari, su una nuova riga, 
 separate da spazi.
*/

package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Split(bufio.ScanWords)
  
  pari := []string{}
  dispari := []string{}

  for scanner.Scan(){
    w := scanner.Text()
    if w == "stop" {
      break
    }

    if len([]rune(w)) % 2 == 0{
      pari = append(pari, w)
    } else {
      dispari = append(dispari, w)
    }
  }

  //riga parole pari
  fmt.Print("Pari: ")
  for i, w := range pari {
    if i > 0 {
      fmt.Print(" ")
    }
    fmt.Print(w)
  }

  fmt.Println()

  //riga parole dispari
  fmt.Print("Dispari: ")
  for i, w := range dispari {
    if i > 0 {
      fmt.Print(" ")
    }
    fmt.Print(w)
  }
}
