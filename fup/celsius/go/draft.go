package main

import "fmt"

func main() {
	var fahrenheit float64
	var celsius float64
	//entrada
	fmt.Scan(&celsius)
	//calculo
	fahrenheit = 1.8*celsius + 32
	//saida
	fmt.Printf("%.6f\n", fahrenheit)

}
