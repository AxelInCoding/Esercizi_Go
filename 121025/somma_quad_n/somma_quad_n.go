package main
import "fmt"

func main(){
  var numero, i, somma int;
  fmt.Print("Inserisci un numero: ");
  fmt.Scan(&numero);

  for i <= numero {
    somma+= i * i;
    fmt.Println("i: ", i, "i*i: ", i * i, "somma :", somma);
    i++;
  }

  fmt.Println("la somma di tutti i quadrati da 1 a n è: ", somma);
}
