/*che legge da standard input una stringa 
e la analizza runa per runa: stampa "nessuna cifra" se la stringa 
non contiene nessuna cifra, stampa invece "ha cifre" altrimenti.*/

package main
import "fmt"

func main () {
  var s string;

  fmt.Print("Inserisci una stringa: ")
  fmt.Scan(&s)

  trovaCifre := false;

  
  for _, r := range s {
    if r >= '0' && r <= '9' {
      trovaCifre = true;
      break; // basta trovare una cifra
    } 
  }

  if trovaCifre {
    fmt.Println("ha cifre")
  } else {
    fmt.Println("nessuna cifra")
  }
}
