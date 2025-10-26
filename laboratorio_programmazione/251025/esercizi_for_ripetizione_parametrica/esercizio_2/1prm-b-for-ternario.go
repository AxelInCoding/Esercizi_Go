package main
import "fmt"
func main() {
    /*
      Questo codice richiede un inserimento da tastiera, poi viene fatto un ciclo for
      che cicli in modo decrescente in base al numero inserito, e viene stampato con gli spazi il numero i-esimo
      andando a capo alla fine del ciclo
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := n; i >= 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

