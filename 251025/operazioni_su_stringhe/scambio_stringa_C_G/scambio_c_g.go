// scambio sulla stringa inserita il carattera 'c' con il carattere 'g' 

package main
import "fmt"


func main (){
  var parola string;
  fmt.Scan(&parola);

  for i:=0; i<len(parola); i++ {
  
    if parola[i] == 'c' {
      fmt.Print("g");
    } else {
      fmt.Print(string(parola[i]));
    }
  }
}
