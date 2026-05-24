package main

import "fmt"

// função que inverte o número
func inverter(n int) int {
	inv := 0

	for n > 0 {
		digito := n % 10
		inv = inv*10 + digito
		n = n / 10
	}

	return inv
}

func main() {
	var num int
	fmt.Scan(&num)

	if num == inverter(num) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
