package main
import (
    "fmt"
    "math/rand"
    "time"
)
func main() {
    /*
    Questo programma genera dei numeri casuali da 0 a 20, per "infinite" volte.
    questi valori vengono stampati, e il ciclo finisce quando il valore generato è uguale a 0.
    la variabile c serve per contare i numeri generati.
    */
    //rand.Seed(time.Now().Unix())
    const K = 20
    var n int

    contatore := 0
    for {
        n = rand.Intn(K+1)
        fmt.Print(n, " ")
        if n == 0 {
            break
        }
        contatore++
    }
    fmt.Println(contatore)
}

