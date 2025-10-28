package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
    /*
    Il programma genera numeri casuali compresi tra 1 e MAX (10) e li stampa su una riga.
La generazione prosegue finché la somma cumulativa dei numeri generati non raggiunge almeno TARGET (20).
Infine stampa la somma totale ottenuta.
    */
    const TARGET = 20
    const MAX = 10
    //rand.Seed(time.Now().Unix())
    var n int

    somma := 0
    for somma < TARGET {
        n = rand.Intn(MAX) + 1
        fmt.Print(n, " ")
        somma += n
    }
    fmt.Print("\n",somma,"\n")
}

