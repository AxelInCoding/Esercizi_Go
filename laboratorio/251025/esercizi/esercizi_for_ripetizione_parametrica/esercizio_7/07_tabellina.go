package main 
import "fmt"

func main (){
  const k = 10;
  var n int;
  fmt.Print("Inserisci un numero per stamparne la tabellina: ");
  fmt.Scan(&n);

  for i:=1; i<=k; i++ {
    fmt.Print(n*i, "\t");
  }
  fmt.Println();
}
