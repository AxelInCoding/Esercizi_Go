/*
Leggere da standard input una sequenza di "parole", 
fino a incontrare EOF (ctrl D su una riga nuova), 
e stampare solo le parole di lunghezza maggiore 
della media delle lunghezze di tutte le parole del 
testo, separate da uno spazio.
*/

package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    
    words := []string{}
    totLen := 0

    for scanner.Scan() {
        w := scanner.Text()
        words = append(words, w)
        totLen += len([]rune(w))
    }

    if len(words) == 0 {
        return
    }

    media := float64(totLen) / float64(len(words))

    first := true
    for _, w := range words {
        if float64(len([]rune(w))) > media {
            if !first {
                fmt.Print(" ")
            }
            fmt.Print(w)
            first = false
        }
    }
}

