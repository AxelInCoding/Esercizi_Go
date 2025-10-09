package main
import (
  "fmt"
  "math/rand"
)

func main (){
  r :=rand.New(rand.NewSource(10)) // crea una sorgente di num pseudo-casuali e con rand.New() crea un nuovo generatore di numeri casuali usando la sorgente dei num casuali

  num :=r.Intn(101); // r è un contenitore del numero casuale estratto

  fmt.Println("Numero casuale generato: ", num);
}
