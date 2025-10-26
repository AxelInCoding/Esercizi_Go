package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; ; i++ {
		quadrato := i * i
		if quadrato > n {
			break
		}
		fmt.Println(quadrato)
	}
}
