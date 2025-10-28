package main
import "fmt"

func main() {
  var input_float float64;
  fmt.Print("Ciao Biondo! Inserisci il tuo numero decimale: ");
  fmt.Scan(&input_float);
  var int_input int = int(input_float);
  fmt.Println(input_float - float64(int_input));
}
