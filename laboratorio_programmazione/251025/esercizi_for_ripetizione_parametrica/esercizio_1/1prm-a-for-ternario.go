package main
import "fmt"
func main() {
    /*
      In questo codice viene fatta una richiesta di inserimento di un numero
      poi viene fatto un ciclo for per iterare lo stesso numero di volte in base al numero inserito.
      Facendo si che l'aggiornamento aumenti ogni volta di +2 ad ogni interazione.
      ed infine viene stampato il contatore del ciclo andando a capo a fine riga.
    */
	var n int
	fmt.Print("un numero: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i = i + 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

