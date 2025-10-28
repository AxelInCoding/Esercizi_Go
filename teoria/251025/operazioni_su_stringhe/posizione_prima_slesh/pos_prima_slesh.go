//trovare la posizione della pirma "/"


package main
import "fmt"

func main (){
  var parola string;
  fmt.Print("Inserisci una parola: ");
  fmt.Scan(&parola);

  var posizione int = -1; // inizializzo a -1 perchè tutte le stringhe incominciano da 0

  for i:=0; i<len(parola); i++{
    if parola[i] == '/' {
      posizione = i;
      break;
    }
  }

  if posizione != -1 {
    fmt.Println("la prima '/' si trova alla posizione: ", posizione)
  } else {
    fmt.Println("Nessuna '/' trovata nella stringa");
  }
} 
