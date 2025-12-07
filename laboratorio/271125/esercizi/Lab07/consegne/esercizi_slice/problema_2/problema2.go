/*
Leggere da standard input una 
sequenza di "parole", fino a incontrare 
la parola "stop", e stampare una parola per riga, 
partendo dall'ultima ("stop" esclusa).
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords) //scanziona per parole, NON per righe

	words := []string{}

	for scanner.Scan() {
		w := scanner.Text()
		if w == "stop" {
			break
		}
		words = append(words, w)
	}

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Println(words[i])
	}
}
