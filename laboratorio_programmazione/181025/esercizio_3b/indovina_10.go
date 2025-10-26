package main
import "fmt"

func main (){
  const NUM_FISSO = 5;
  var numero int;
  
  fmt.Print("Inserisci un numero tra 1 e 10: ");
  fmt.Scan(&numero);

  if numero <= 1 || numero >= 10 {
    fmt.Println("Valore non valido");
  }else if numero == NUM_FISSO {
    fmt.Println("Hai indovinato!");
  }else{
    fmt.Println("Non hai indovinato");
  }
}
