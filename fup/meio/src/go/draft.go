package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Scanln(&num3)

	var mediano int
	if (num1 > num2 && num1 < num3) || (num1 > num3 && num1 < num2) {
		mediano = num1
	} else if (num2 > num1 && num2 < num3) || (num2 > num3 && num2 < num1) {
		mediano = num2
	} else {
		mediano = num3
	}
	fmt.Println(mediano)
}
