package main
import (
	"fmt"
	"io"
)

func main() {
	var parola string
	var minParola string
	var pos, minPos int
	var first bool = true

	for {
		_, err := fmt.Scan(&parola)
		if err == io.EOF {
			break
		}
		pos++
		if first {
			minParola = parola
			minPos = pos
			first = false
		} else if parola < minParola {
			minParola = parola
			minPos = pos
		}
	}

	if !first { // cioè: se è stata letta almeno una parola
		fmt.Println(minParola, minPos)
	}
}

