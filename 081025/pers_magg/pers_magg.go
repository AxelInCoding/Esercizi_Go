package main
import "fmt"

func main (){
  var annoNascita int;
  var annoCorrente int;
  
  fmt.Print("Inserisci il tuo anno di nascita: ");
  fmt.Scan(&annoNascita);

  fmt.Print("Inserisci l'anno corrente: ");
  fmt.Scan(&annoCorrente);
  
  if annoNascita > 0 && annoNascita != 0 {
     eta := annoCorrente - annoNascita;
     
    if eta >= 18 {
      fmt.Println("Sei maggiorenne!");
    }else if annoNascita > annoCorrente {
      fmt.Println("L'anno di nascita non può essere maggiore dell'anno corrente!");
    }else{
      fmt.Println("NON sei maggiorenne!");
    }
    
  }else{
    fmt.Println("Errore L'anno di nascita deve essere un numero positivo!");
  }
}

