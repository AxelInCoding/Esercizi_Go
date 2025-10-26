package main 
import "fmt"

func main(){
  const FATTORE = 9.0 / 5.0 // 1 grado C = 9/5 F
  const OFFSET = 32.0

  var gradiCentigradi float64;

  fmt.Print("Inserisci la temperature in gradi Cnetigradi: ");
  _, err := fmt.Scan(&gradiCentigradi);
  if err != nil {
    fmt.Println("Errore nella temperatura del valore: ", err);
  }

  gradiFahrenheit :=gradiCentigradi * FATTORE + OFFSET;

  fmt.Println(gradiCentigradi,"C = ",  gradiFahrenheit, "F");
}
