package main 
import "fmt"

/*

In un file operazioni.go definire due versioni di una stessa funzione operazioni, 
una versione con variabili di ritorno nominate (operazioniV1) e una senza (operazioniV1), 
che accettano due interi e restituiscono somma, prodotto e differenza in quest'ordine:

- operazioniV1(n1, n2 int) (int, int, int) //con variabili di ritorno nominate

- operazioniV2(n1, n2 int) (int, int, int) //senza variabili di ritorno nominate

Scrivere un main che legge da standard input due int e invoca le due funzioni per testarle.

nomefile: operazioni.go

*/

//Variabili di ritorno nominate
func operazioniV1(n1, n2 int) (somma int, prodotto int, differenza int) {
  somma = n1 + n2
  prodotto = n1 * n2
  differenza = n1 - n2
  return
}

//Variabili di ritorno NON nominate
func operazioniV2(n1, n2 int) (int, int, int) {
  return n1 + n2, n1 * n2, n1 - n2
}

func main () {
  var a, b int
	
	fmt.Print("Inserisci due numeri: ")
	fmt.Scan(&a, &b)

	// Versione con variabili di ritorno nominate
	s1, p1, d1 := operazioniV1(a, b)
	fmt.Println("operazioniV1:")
	fmt.Println("Somma:", s1, "Prodotto:", p1, "Differenza:", d1)

	// Versione senza variabili di ritorno nominate
	s2, p2, d2 := operazioniV2(a, b)
	fmt.Println("operazioniV2:")
	fmt.Println("Somma:", s2, "Prodotto:", p2, "Differenza:", d2)

}
