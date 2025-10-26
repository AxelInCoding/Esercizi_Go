/*
1. Quante diverse aliquote da applicare ci sono?
   - Ci sono 2 aliquote: 10% e 24%.

2. Quanti if/else occorrono come minimo? Di che tipo? Perché?
   - Occorrono almeno 2 if/else a più vie:
     1. Uno per distinguere se l’utente è coniugato o non coniugato.
     2. Uno per distinguere se il reddito rientra nello scaglione basso o alto.
   - Entrambi devono essere if/else perché le condizioni sono mutualmente esclusive:
     - reddito <= scaglione → aliquota bassa
     - reddito > scaglione → aliquota alta
     - coniugato o non coniugato influenza lo scaglione da considerare.
*/


package main

import "fmt"

func main() {
	const (
		ALIQUOTA_BASSA = 10
		ALIQUOTA_ALTA  = 24
		SCAGLIONE_NC   = 32000 // non coniugato
		SCAGLIONE_C    = 64000 // coniugato
	)

	var reddito int
	var coniugato bool

	fmt.Print("reddito? ")
	fmt.Scan(&reddito)
	fmt.Print("coniugato? (true/false) ")
	fmt.Scan(&coniugato)

	var tasse int

	if coniugato {
		if reddito <= SCAGLIONE_C {
			tasse = reddito * ALIQUOTA_BASSA / 100
		} else {
			tasse = reddito * ALIQUOTA_ALTA / 100
		}
		fmt.Print("tasse per coniugato con reddito ", reddito, ": ", tasse)
	} else {
		if reddito <= SCAGLIONE_NC {
			tasse = reddito * ALIQUOTA_BASSA / 100
		} else {
			tasse = reddito * ALIQUOTA_ALTA / 100
		}
		fmt.Print("tasse per non coniugato con reddito ", reddito, ": ", tasse)
	}
}

