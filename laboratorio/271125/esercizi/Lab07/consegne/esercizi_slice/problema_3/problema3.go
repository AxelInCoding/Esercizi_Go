/*
Leggere da standard input una sequenza di "parole",  
fino a incontrare la parola "stop", e stampare
 solo le parole di lunghezza pari 
 ("stop" esclusa), una parola per riga.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
  scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) //scansione per parole, NON per righe

	for scanner.Scan() {
		w := scanner.Text()
		if w == "stop" {
			break
		}
		
		if len([]rune(w)) % 2 == 0 {
      fmt.Println(w)
		}
	}
}
