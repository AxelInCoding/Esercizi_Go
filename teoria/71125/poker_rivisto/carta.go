package main 

type carta struct{
  seme int //0->3 con importanza crescente
  valore int //0->12 con importanza crescente (12=A)
}

func String(x carta) string{
  nomeSeme := [...]string{"\u2664", "\u26667","\u2662","\u2661"}
  nomeValore := [...]string{"2","3","4","5", "6","7", "8", "9", "10", "J","Q", "K", "A"}

  return nomeValore[x.valore] + nomeSeme[x.seme]
}


