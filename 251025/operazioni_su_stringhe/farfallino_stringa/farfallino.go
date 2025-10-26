// da CANE --> CAFANEFE

package main
import "fmt"


func main (){
  var parola string;
  fmt.Scan(&parola);

  for i:=0; i< len(parola); i++ {
  fmt.Print(string(parola[i]));
  
    if parola[i] == 'a' || parola[i] == 'e' || parola[i] == 'i' || parola[i] == 'o' || parola[i] == 'u' {
      fmt.Print("f", string(parola[i]));
    }
  }
}
