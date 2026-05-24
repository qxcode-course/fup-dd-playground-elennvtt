package main

import "fmt"

func main() {
	var N int
	var pe string

	fmt.Scan(&N)
	fmt.Scan(&pe)

	atual := pe
	primeiro := true

	fmt.Print("[ ")

	for i := 0; i <= 10; i++ {

		if i == N {
			continue
		}

		// céu
		if i == 10 {
			if N != 10 {
				if !primeiro {
					fmt.Print(" ")
				}
				fmt.Print("ceu")
			}
			break
		}

		if !primeiro {
			fmt.Print(" ")
		}

		fmt.Printf("%d%s", i, atual)
		primeiro = false

		// alterna o pé
		if atual == "d" {
			atual = "e"
		} else {
			atual = "d"
		}
	}

	fmt.Print(" ]")
}