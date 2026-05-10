package main

import "fmt"

func main() {
	var nota1, nota2, notafinal int

	fmt.Scan(&nota1)
	fmt.Scan(&nota2)
	fmt.Scan(&notafinal)

	media := float64(nota1+nota2) / 2.0

	if media >= 7 {
		fmt.Println("aprovado")
	} else if media < 4 {
		fmt.Println("reprovado")
	} else {
		novamedia := (media + float64(notafinal)) / 2.0

		if novamedia >= 5 {
			fmt.Println("aprovado na final")
		} else {
			fmt.Println("reprovado na final")
		}
	}
}
