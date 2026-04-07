package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, y1, y2 float64

	//entrda
	fmt.Scan(&x1)
	fmt.Scan(&y1)
	fmt.Scan(&x2)
	fmt.Scan(&y2)

	// distancia
	distancia := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	//saida
	fmt.Printf("%.2f\n", distancia)

}
