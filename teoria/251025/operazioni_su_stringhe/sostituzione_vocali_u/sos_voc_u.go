// sostituzione di tutte le vocali con il carattere 'u'

package main 
import "fmt"

func main (){
  var parola string;
  fmt.Scan(&parola);

  for _, r := range parola {
    if r == 'a' || r == 'e' || r == 'i' || r == 'o' {
      fmt.Print("u");
    } else {
      fmt.Print(string(r));
    }
  }
}
