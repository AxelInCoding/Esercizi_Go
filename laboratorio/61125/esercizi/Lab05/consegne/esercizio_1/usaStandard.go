package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const VOCALI = "aeiouAEIOU"
	const s1 = "c"
	const s2 = "sc"

	var parola string
	fmt.Print("Inserisci una parola: ")
	fmt.Scan(&parola)

	if strings.Contains(parola, s2) {
		fmt.Println("(1)", parola, "contiene", s2)
	} else {
		fmt.Println("(1)", parola, "non contiene", s2)
	}

	if strings.ContainsAny(parola, VOCALI) {
		fmt.Println("(2)", parola, "ha vocali")
	} else {
		fmt.Println("(2)", parola, "non ha vocali")
	}

	fmt.Println("(3)", parola, "contiene", strings.Count(parola, s1), s1)

	// trova la prima e ultima vocale
	prima := -1
	ultima := -1
	for i, c := range parola {
		if strings.ContainsRune(VOCALI, c) {
			if prima == -1 {
				prima = i
			}
			ultima = i
		}
	}

	if prima != -1 {
		fmt.Println("(4) la prima vocale di", parola, "è in posizione", prima, ", l'ultima in posizione", ultima)
	} else {
		fmt.Println("(4)", parola, "non contiene vocali")
	}

	fmt.Println("(5)", strings.Repeat(s2, 3))
	fmt.Println("(6)", strings.Repeat(s1, 5))

	s := ""
	if strings.ContainsAny(parola, "0123456789") {
		for i := 0; i < len(parola); i++ {
			if strings.ContainsAny(string(parola[i]), "0123456789") {
				s += string(parola[i])
			}
		}
		fmt.Println("(7) stringa", s)

		n, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			fmt.Printf("(8) intero %d\n", n)
		} else {
			fmt.Println("(8) errore nella conversione:", err)
		}
	}

	if s != "" {
		fstr := "0." + s
		f, err := strconv.ParseFloat(fstr, 64)
		if err == nil {
			fmt.Printf("(9) float %f\n", f)
		} else {
			fmt.Println("(9) errore nella conversione a float:", err)
		}
	}
}

