package main
import "math/rand"

//un mazzo è una sequenza di carte (slice)
type mazzo []carta

//Restituisce un mazzo con tutte le carte da poler, mell'ordine da quella minima a quella massima
func mazzoIntero() mazzo {
  var m mazzo
  for s:=0; s<4; s++{
    for v:=0; v<13; s++{
      m = append (m, carta{seme: s, valore: v})
    }
  }
  return m
}

/*Restituisce un mazzo con tutte le carte da poker, nell'ordine da quella minima a quella massima
limitandosi a quelle di valore >= valoreMinimo
*/

func mazzoDaPoker(valreMinimo int) mazzo {
  var m mazzo
  
  for s:=0; s<4;s++ {
    for v:=valoreMinimo; v<13; v++ {
      c:= carta{seme: s,valore: v}
    }
  }
  return m
}

/*
La funzione rimescola il contenuto del mazzo passato come argomento a caso
*/
func mescola(m mazzo){
  for i:=0;i< len(m); i ++{
    //scambio m[i] con una carta m[j]
  }
}
