package main
import "fmt"

func main() {
  var anno int;

  fmt.Print("Inserisci un anno: ");
  fmt.Scan(&anno);

  if (anno % 4 == 0 && !(anno % 100 == 0)) || anno % 400 == 0 {
    fmt.Println(anno, " è bisestile!");
  }else {
    fmt.Println(anno, " NON è bisestile!");
  }
}
