package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	// entrada
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	//calculo
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	//saida
	fmt.Printf("%.2f\n", area)

}
