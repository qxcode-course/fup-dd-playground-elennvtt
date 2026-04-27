package main

import "fmt"

func main() {
	var a, b int
	var operador string

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&operador)

	var resultado int

	switch operador {
	case "+":
		resultado = (a + b)
	case "-":
		resultado = (a - b)
	case "*":
		resultado = (a * b)
	case "/":
		resultado = (a / b)
	default:
		fmt.Println("ERROR")
		return

	}
	fmt.Println(resultado)

}
