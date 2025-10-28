package main

import "fmt"

func main() {
	var num1, den1, num2, den2 int

	// Input prima frazione
	fmt.Print("num e den fraz 1: ")
	fmt.Scan(&num1, &den1)

	// Input seconda frazione
	fmt.Print("num e den fraz 2: ")
	fmt.Scan(&num2, &den2)

	if num1*den2 == num2*den1 {
		fmt.Println("equivalenti")
	} else {
		fmt.Println("non equivalenti")
	}
}

