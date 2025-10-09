/*
Problema: Scrivere un programma Go scambio.go che, dati in input due valori di tipo int,
controlli che i valori letti siano effettivamente di tipo int (come fare? vedere la documentazione di fmt.Scan), li stampi nell’ordine in cui sono stati forniti,
scambi i due valori (senza appoggiarsi a altre variabili), e poi esegua la stessa istruzione di
stampa (per mostrare che i valori sono stati effettivamente scambiati tra le due var).
Provare il programma su input interi, float, testo, e osservare che risultati si ottengono.
*/

package main
import "fmt"

func main(){
  var valore1, valore2 int;

  fmt.Print("Inserisci i valori (separali da uno spazio): ");
  n,err := fmt.Scan(&valore1,&valore2); //questa riga controlla se i valori inseriti dall'utente sono dello stesso tipo in cui sono stati dichiarati nel programma (int)
  
  if err != nil || n!=2 {
    fmt.Println("Errore: assicurati di inserire due numeri interi");
  }

  fmt.Println("Stampa dei valori nell'ordine in cui sono stati inseriti: ", valore1, valore2);

  valore1, valore2 = valore2, valore1; //scambio dei valori
  
  fmt.Println("Stampa dei valori nell'ordine scambiato :", valore1, valore2);
}
