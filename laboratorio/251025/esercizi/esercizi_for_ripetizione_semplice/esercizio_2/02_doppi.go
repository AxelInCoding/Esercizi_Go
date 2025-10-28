package main
import "fmt"

func main (){
  const k=5;
  var numero float64;

  fmt.Println("Inserisci ", k, " numeri: ");

  for i:=0; i<k; i++ {
    fmt.Print("Numero: ");
    fmt.Scan(&numero);
    fmt.Println("Il doppio è:", numero*2);
  }
}
