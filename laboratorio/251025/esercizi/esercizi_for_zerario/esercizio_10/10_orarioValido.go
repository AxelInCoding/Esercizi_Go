package main 
import "fmt"

func main (){
  var ore, minuti int;

  for {
    fmt.Print("ore e minuti: ");
    fmt.Scan(&ore, &minuti);

    if (ore >=0 && ore <=23) && (minuti >=0 && minuti <=59) {
      fmt.Println("orario valido");
      break;
    }
  }
}
