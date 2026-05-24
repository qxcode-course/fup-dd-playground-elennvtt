package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	fmt.Print("[ ")

	for i := A; i <= B; i++ {
		fmt.Printf("%d %d", i, B-(i-A))

		if i != B {
			fmt.Print(" ")
		}
	}

	fmt.Println(" ]")
}