package main

import "fmt"

func main() {
	var tempo int
	//entrada
	fmt.Scan(&tempo)
	//calculo horas
	horas := tempo / 3600
	//resto horas 
	resto := tempo % 3600
	// minutos
	minutos := resto / 60
	//segundos
	segundos := resto % 60

	// saida 
	fmt.Printf("%d:%d:%d\n", horas, minutos, segundos)


}
